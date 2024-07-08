package main_database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Conection() {
	err := godotenv.Load()
	config := mysql.Config{}
	db, err := sql.Open("mysql", config.FormatDSN())
}
