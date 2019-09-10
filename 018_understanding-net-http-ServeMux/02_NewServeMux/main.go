package main

import (
	"io"
	"net/http"
)

type falooda int 

func (m falooda) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I am Falooda")
}

type kulfi int

func (k kulfi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I am Kulfi")
}

func main() {
	var d falooda
	var e kulfi
	mux := http.NewServeMux()
	mux.Handle("/falooda", d)
	mux.Handle("/kulfi", e)
	http.ListenAndServe(":8080", mux)
}

