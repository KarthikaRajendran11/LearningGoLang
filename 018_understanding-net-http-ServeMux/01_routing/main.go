package main

import (
	"io"
	"net/http"
)

type falooda int

func (m falooda) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch(req.URL.Path) {
	case "/go":
		io.WriteString(w, "GoLang")
	case "/node":
		io.WriteString(w, "NodeJS")
	}
}

func main() {
	var d falooda
	http.ListenAndServe(":8080", d)
}