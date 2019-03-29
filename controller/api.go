package controller

import (
	"crud/model"
	"crud/service"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

func Employee(w http.ResponseWriter, r *http.Request) {
	logs.Info("METHOD : ", r.Method)

	if r.Method == "GET" {
		id := r.URL.Query().Get("id")

		if id != "" {
			intId, err := strconv.Atoi(id)

			if err != nil {
				panic(err.Error())
			}

			doGetById(w, r, intId)
		} else {
			doGet(w, r)
		}
	} else if r.Method == "POST" {
		doPost(w, r)
	} else if r.Method == "PUT" {
		doPut(w, r)
	} else if r.Method == "DELETE" {
		doDelete(w, r)
	}
}

func doGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employees := service.FindAllEmployee()

	response := model.Response{}
	response.Status = 200
	response.Data = employees

	json.NewEncoder(w).Encode(response)
}

func doGetById(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")

	employees := service.FindEmployeeById(id)

	response := model.Response{}
	response.Status = 200
	response.Data = employees

	json.NewEncoder(w).Encode(response)
}

func doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := r.FormValue("name")
	city := r.FormValue("city")

	employee := service.SaveEmployee(name, city)

	response := model.Response{}
	response.Status = 201
	response.Data = employee

	json.NewEncoder(w).Encode(response)
}

func doPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		panic(err.Error())
	}

	name := r.FormValue("name")
	city := r.FormValue("city")

	employee := service.UpdateEmployee(name, city, intId)

	response := model.Response{}
	response.Status = 200
	response.Data = employee

	json.NewEncoder(w).Encode(response)
}

func doDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	service.DeleteEmployee(intId)

	response := model.Response{}
	response.Status = 200
	response.Data = "Deleted"

	json.NewEncoder(w).Encode(response)
}
