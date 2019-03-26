package main

import (
	"crud/controller"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:8080")

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/update", controller.Update)
	http.HandleFunc("/delete", controller.Delete)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
}
