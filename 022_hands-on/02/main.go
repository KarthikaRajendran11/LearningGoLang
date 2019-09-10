package main

import (
	"net/http"
	"html/template"
)
func a(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "I love food")
}
func b(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "I love idilies")
}
func c(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "I love dosas")
}
func d(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "I am Karthika Rajendran")
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("/Users/karthika.rajendran/Documents/repo/golang/022_hands-on/02/index.html"))
}
func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/idly", b)
	http.HandleFunc("/dosa", c)
	http.HandleFunc("/me", d)
	http.ListenAndServe(":8080", nil)
}