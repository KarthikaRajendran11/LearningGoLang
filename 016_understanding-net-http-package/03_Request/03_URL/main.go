package main

import (
	"net/http"
	"net/url"
	"html/template"
	"log"
)

type falooda int

func (m falooda) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method string
		URL *url.URL
		Submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.html", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("/Users/karthika.rajendran/Documents/repo/golang/016_understanding-net-http-package/03_Request/03_URL/index.html")) 
}

func main() {
	var d falooda 
	http.ListenAndServe(":8080", d)
}