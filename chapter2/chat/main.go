package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
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
	t.tmpl.Execute(w, r)
}

func main() {
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
