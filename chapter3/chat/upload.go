package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid") // htmlのformにhiddenで渡された値を取得
	file, header, err := req.FormFile("avatarFile")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	filename := filepath.Join("avatars", userId+filepath.Ext(header.Filename))
	err = os.WriteFile(filename, data, 0777)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	// io.WriteString(w, "成功")
	// io.WriteString(w, "3秒後にtopページにリダイレクトします")

	// io.WriteString, http.Redirect, http.WriteHeader()は全て、response.WriteHeaderのインターフェースを持つので、
	// 複数回呼び出すと2個目以降の呼び出しは実行されない。
	http.Redirect(w, req, "/chat", http.StatusTemporaryRedirect)

}
