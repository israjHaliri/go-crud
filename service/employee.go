package service

import (
	"crud/config"
	"crud/model"
	_ "github.com/go-sql-driver/mysql"
)

func FindAllEmployee() []model.Employee {
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

func FindEmployeeById(id int) model.Employee {
	db := config.Connection()

	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", id)

	if err != nil {
		panic(err.Error())
	}

	employee := model.Employee{}

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
	}

	defer db.Close()

	return employee
}

func SaveEmployee(name string, city string) {
	db := config.Connection()

	prepare, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")

	if err != nil {
		panic(err.Error())
	}

	prepare.Exec(name, city)

	defer db.Close()
}

func UpdateEmployee(name string, city string, uid int) {
	db := config.Connection()

	prepare, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	prepare.Exec(name, city, uid)

	defer db.Close()
}

func DeleteEmployee(id int) {
	db := config.Connection()

	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)

	defer db.Close()
}
