package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (db *sql.DB) {
	driver := "mysql"
	url := "root:root@/golang"

	connection, e := sql.Open(driver, url)

	if e != nil {
		panic(e.Error())
	}

	return connection
}
