// server.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	layoutPath := filepath.Join("templates", "layout.html.tmpl")
	pagePath := filepath.Join("templates", tmpl+".html.tmpl")

	tmpls, err := template.ParseFiles(layoutPath, pagePath)
	if err != nil {
		http.Error(w, "Unable to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpls.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template: "+err.Error(), http.StatusInternalServerError)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Go, HTMX, Tailwind CSS",
	}

	renderTemplate(w, "index", data)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Go, HTMX, Tailwind CSS",
	}

	renderTemplate(w, "form", data)
}

// クロージャーを使用して引数を渡す
func testHandler(fileName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmplPath := fmt.Sprintf("templates/%s.html.tmpl", fileName)

		t, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Unable to load template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, "Unable to render template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
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
	http.HandleFunc("/test", testHandler("test"))
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/update", updateHandler)

	log.Println("Server started: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
