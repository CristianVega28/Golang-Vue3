package data

import (
	"cristianvega6150/server/database"
	"fmt"
	"testing"
)

func TestGetDataArray(t *testing.T) {
	data := GetDataJsonClients()

	fmt.Println(data[0])
	test := database.Helper_template_insert(data[0])
	fmt.Println(test)
	if data != nil {
		fmt.Println("Si existe datos")
		// fmt.Println(data[0])
	}
}
