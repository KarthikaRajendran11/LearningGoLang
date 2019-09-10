package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dosa)
	http.HandleFunc("/toby.jpg", toby)
	http.ListenAndServe(":8080", nil)
}

func dosa(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func toby(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}