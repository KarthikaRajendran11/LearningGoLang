package main

import (
	"net/http"
	"io"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/idly", idly)
	http.ListenAndServe(":8080", nil)
}

func idly(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg>`)
}