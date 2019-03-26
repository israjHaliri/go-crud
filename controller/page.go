package controller

import (
	"crud/service"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles(
	"template/layout/Header.tmpl",
	"template/layout/Menu.tmpl",
	"template/layout/Footer.tmpl",
	"template/Index.tmpl",
	"template/New.tmpl",
	"template/Edit.tmpl",
))

func Index(w http.ResponseWriter, r *http.Request) {
	res := service.FindAllEmployee()

	tmpl.ExecuteTemplate(w, "Index", res)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")

		service.SaveEmployee(name, city)

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

	res := service.FindEmployeeById(intId)

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

		service.UpdateEmployee(name, city, intId)

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

	service.DeleteEmployee(intId)

	http.Redirect(w, r, "/", 301)
}
