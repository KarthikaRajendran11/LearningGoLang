package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dosa)
	http.ListenAndServe(":8080", nil)
}

func dosa(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset-utf-8")
	io.WriteString(w, `<!--image doesn't serve-->
	<img src="/toby.jpg">`)
}