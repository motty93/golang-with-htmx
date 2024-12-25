// server.go
package main

import (
	"html/template"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/index.html"))
	tmpl.Execute(w, nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/form.html"))
	tmpl.Execute(w, nil)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<p>これは新しいコンテンツです！</p>"))
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		w.Write([]byte("<p>こんにちは" + name + "</p>"))
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/update", updateHandler)

	http.ListenAndServe(":8080", nil)
}
