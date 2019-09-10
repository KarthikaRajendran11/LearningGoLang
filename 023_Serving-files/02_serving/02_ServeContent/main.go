package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", idly)
	http.HandleFunc("/toby.jpg", chutney)
	http.ListenAndServe(":8080", nil)
}

func idly(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func chutney(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File not found", 404)
		return 
	}
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}