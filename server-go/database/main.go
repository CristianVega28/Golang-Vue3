package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type ServiceDB struct {
}

func (*ServiceDB) Conection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error en las variables de entorno")
	}
	address := [3]string{os.Getenv("DB_HOST"), ":", os.Getenv("DB_PORT")}
	config := mysql.Config{
		DBName:               os.Getenv("DB_DATABASE"),
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 strings.Join(address[:], ""),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())

	pingError := db.Ping()
	if pingError != nil {
		log.Fatal(pingError.Error())
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
