package controller

import (
	"crud/service"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := service.FindEmployeeById(1)
	json.NewEncoder(w).Encode(res)
}
