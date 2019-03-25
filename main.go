package main

import (
	"crud/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("template/*"))

func Newndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := service.FindById(1)
	json.NewEncoder(w).Encode(res)
}

func Index(w http.ResponseWriter, r *http.Request) {
	res := service.FindAll()

	tmpl.ExecuteTemplate(w, "Index", res)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")

		service.Save(name, city)

		http.Redirect(w, r, "/", 301)
	} else {
		log.Println("Method not post")
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		panic(err.Error())
	}

	res := service.FindById(intId)

	tmpl.ExecuteTemplate(w, "Edit", res)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("id")

		intId, err := strconv.Atoi(id)

		if err != nil {
			panic(err.Error())
		}

		service.Update(name, city, intId)

		http.Redirect(w, r, "/", 301)
	} else {
		log.Println("Method not post")
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		panic(err.Error())
	}

	service.Delete(intId)

	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/newdwx", Newndex)

	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
}
