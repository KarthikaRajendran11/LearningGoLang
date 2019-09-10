package main

import (
	"net/http"
	"io"
	"html/template"
)

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func idly(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "/Users/karthika.rajendran/Documents/repo/golang/024_hands-on/01/toby.jpg")
	tpl.ExecuteTemplate(w, "idly.gohtml", "This is idly")
}

// func dosa(w http.ResponseWriter, req *http.Request) {
// 	http.ServeFile(w, req, "/Users/karthika.rajendran/Documents/repo/golang/024_hands-on/01/toby.jpg")
// }

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("/Users/karthika.rajendran/Documents/repo/golang/024_hands-on/01/idly.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/idly/", idly)
	// http.HandleFunc("/idly.jpg", dosa)
	http.ListenAndServe(":8080", nil)
}
