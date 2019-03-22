package service

import (
	"crud/config"
	_ "github.com/go-sql-driver/mysql"
	"crud/model"
)

func FindAll() []model.Employee {
	db := config.Connection()

	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")

	if err != nil {
		panic(err.Error())
	}

	employee := model.Employee{}
	result := []model.Employee{}

	for selDB.Next() {
		var id int
		var name, city string

		err = selDB.Scan(&id, &name, &city)

		if err != nil {
			panic(err.Error())
		}

		employee.Id = id
		employee.Name = name
		employee.City = city

		result = append(result, employee)
	}

	defer db.Close()

	return result
}
