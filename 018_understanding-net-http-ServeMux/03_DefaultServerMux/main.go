package main

import (
	"io"
	"net/http"
)

func a(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "func 1")
}

func b(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "func2")
}

func main() {
	http.HandleFunc("/falooda", a)
	http.HandleFunc("/nodejs", b)
	http.ListenAndServe(":8080", nil)
}