package main

import (
	"html/template"
	"net/http"
	"net/url"
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
		Submissions map[string][]string
		Header http.Header
		Host string 
		ContentLength int64 
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index.html", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("/Users/karthika.rajendran/Documents/repo/golang/016_understanding-net-http-package/03_Request/05_Host_ContentLength/index.html"))
}

func main() {
	var d falooda
	http.ListenAndServe(":8080", d)
}