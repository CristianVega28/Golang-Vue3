package seeders

import (
	"cristianvega6150/server/database"
	"database/sql"
	"fmt"
)

func ClientSeedStart(decode_data []database.Clientes, connection *sql.DB) {

	var query string = "INSERT INTO clientes VALUES "

	result := database.Helper_insert_data_template(decode_data)

	query += result

	fmt.Println(query)
	smt, err := connection.Prepare(query)

	if err != nil {
		fmt.Println(err.Error())
	}
	// defer smt.Close()

	smt.Exec()

}
