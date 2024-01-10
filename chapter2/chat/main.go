package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"github.com/tker-78/blueprint/chat/config"
)

type templateHandler struct {
	once     sync.Once
	filename string
	tmpl     *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.tmpl.Execute(w, data)
}

func main() {

	gomniauth.SetSecurityKey(config.Google.SecurityKey)
	gomniauth.WithProviders(
		google.New(config.Google.ClientId, config.Google.ClientSecret, config.Google.URL),
	)

	r := newRoom() // roomを生成

	mux := http.NewServeMux()
	mux.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	mux.Handle("/login", &templateHandler{filename: "login.html"})
	mux.HandleFunc("/auth/", loginHandler)
	mux.Handle("/room", r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go r.run() // roomを起動

	server.ListenAndServe()

}
