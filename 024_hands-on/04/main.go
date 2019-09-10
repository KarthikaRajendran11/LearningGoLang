package main

import (
	"net/http"
	"log"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("/Users/karthika.rajendran/Documents/repo/golang/024_hands-on/04/starting-files/templates/index.gohtml"))
}

func idly(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", idly)
	http.Handle("/resources", http.StripPrefix("/resources", http.FileServer(http.Dir("/Users/karthika.rajendran/Documents/repo/golang/024_hands-on/04/starting-files/public/pics"))))
	http.ListenAndServe(":8080", nil)
}