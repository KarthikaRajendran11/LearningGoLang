package main

import (
	"fmt"
	"net/http"
)

type falooda int

func (m falooda) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Mcleod-Key", "this is from McLeod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var d falooda
	http.ListenAndServe(":8080", d)
}