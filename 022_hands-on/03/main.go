package main

import (
	"net/http"
	"html/template"
)

type falooda int 

func (a falooda) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "I love food ")
}

type kulfi int

func (b kulfi) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "I love idlies")
}

type rasgulla int

func (c rasgulla) ServeHTPP(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "I love dosas")
}

type gulabjamun int

func (d gulabjamun) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "I am Karthika Rajendran")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("/Users/karthika.rajendran/Documents/repo/golang/022_hands-on/03/index.html"))
}

func main() {
	var a falooda
	var b kulfi
	// var c rasgulla 
	var d gulabjamun
	mux := http.NewServeMux()
	http.Handle("/", a)
	http.Handle("/idly", b)
	// http.Handle("/dosa", c) 
	http.Handle("/me", d) 
	http.ListenAndServe(":8080", mux)
}