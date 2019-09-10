package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/Users/karthika.rajendran/Documents/repo/golang/024_hands-on/03/starting-files/public/pics")))
	http.ListenAndServe(":8080", nil)
}