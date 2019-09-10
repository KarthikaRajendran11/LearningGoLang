package main

import (
	"net/http"
	"log"
	"html/template"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.ListenAndServe(":8080", nil)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("/Users/karthika.rajendran/Documents/repo/golang/024_hands-on/05/templates/*"))
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.html", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		HandleError(w, err)
		return
	}
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)


func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}