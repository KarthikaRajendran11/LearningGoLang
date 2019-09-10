package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", idly)
	http.HandleFunc("/toby.jpg", toby)
	http.ListenAndServe(":8080", nil)
}

func idly(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">`)
}

func toby(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpg")
}