package main

import (
	"html/template"
	"log"
	"net/http"
)

type falooda int

func (m falooda) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.html", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("/Users/karthika.rajendran/Documents/repo/golang/016_understanding-net-http-package/03_Request/01_ParseForm_Form/index.html"))
}

func main() {
	var d falooda
	http.ListenAndServe(":8081", d)
}
