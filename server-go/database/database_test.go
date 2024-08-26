package database

import (
	// "cristianvega6150/server/database/seeders"

	"fmt"
	"testing"
)

type userPrueba struct {
	Id   int
	Name string
	Age  int
}

func TestTemplateExample(t *testing.T) {

	query := "INSERT INTO prueba VALUES "
	user := []userPrueba{
		userPrueba{
			Id: 1,
			// Name: "cristian",
			Age: 21,
		},
		userPrueba{
			Id:   2,
			Name: "erick",
			Age:  20,
		},
		userPrueba{
			Id:   3,
			Name: "valerie",
			Age:  10,
		},
	}

	message_final := Helper_insert_data_template(user)

	query += message_final

	// Es una query valida en sql para poder insertar datos masivos
	fmt.Println(query)
}
