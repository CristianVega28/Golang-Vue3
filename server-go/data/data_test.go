package data

import (
	"cristianvega6150/server/database"
	"fmt"
	"testing"
)

func TestGetDataArray(t *testing.T) {
	data := GetDataJsonClients()

	test := database.Helper_insert_data_template(data[0:2])
	fmt.Println(test)
	if data != nil {
		fmt.Println("Si existe datos")
		fmt.Println(data[0])
	}
}
