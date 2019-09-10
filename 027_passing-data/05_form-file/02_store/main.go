package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	var tpl *template.Template

	tpl = template.Must(template.ParseGlob("/Users/karthika.rajendran/Documents/repo/golang/027_passing-data/05_form-file/02_store/templates/*"))
	
	if req.Method == http.MethodPost {
		// open a file
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// print some info
		fmt.Println(h)

		// read its content
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		// create another file and write the prev content to the new file
		fmt.Println("File Name : " + h.Filename)
		dst, err := os.Create(filepath.Join("/Users/karthika.rajendran/Documents/repo/golang/027_passing-data/05_form-file/02_store/user/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", s)
}