package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", idly)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func idly(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	fmt.Println(w, "Look for string : " + v)
}