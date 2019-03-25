package main

import (
	"crud/service"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("template/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	res := service.FindAll()

	tmpl.ExecuteTemplate(w, "Index", res)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
}
