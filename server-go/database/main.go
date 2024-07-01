package main_database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func Conection() {
	config := mysql.Config{}
	db, err := sql.Open("mysql", config.FormatDSN())
}
