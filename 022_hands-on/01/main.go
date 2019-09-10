package main

import (
	"io"
	"net/http"
)

func a(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I love food")
}

func b(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I love idlies")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I love dosas")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am Karthika Rajendran")
}

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/idly", b)
	http.HandleFunc("/dosa", c)
	http.HandleFunc("/me", d)
	http.ListenAndServe(":8080", nil)
}