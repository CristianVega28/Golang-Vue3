package data

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

// Struct para representar a los clientes
type Cliente struct {
	ID               string           `json:"customer_id"`
	FullName         string           `json:"full_name"`
	BirthDate        string           `json:"birth_date"`
	LocationCustomer LocationCustomer `json:"location_customer"`
	Phone            int              `json:"phone"`
	Email            string           `json:"email"`
	DischargeDate    string           `json:"discharge_date"`
	CustomerGroup    string           `json:"customer_group"`
}

type LocationCustomer struct {
	CustomerAddress    string `json:"customer_address"`
	CustomerSector     string `json:"customer_sector"`
	CustomerPostalCode string `json:"customer_costal_code"`
}

/*
Está función es para poder sacar información dentro del excel, con estos datos luego
*/

func GetData(index_data int) {

	// var data []Cliente
	currentDir, errPath := os.Getwd()
	file, err := excelize.OpenFile(filepath.Join(currentDir, "\\data\\clientes.xlsx"))

	if err != nil || errPath != nil {
		fmt.Println("error en las rutas")
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := file.GetRows("Hoja1")

	if err != nil {
		log.Fatal("Error al Abria la hoja del excel")
	}
	for index, row := range rows {

		/*
			|---ID---|---FullName---|---BirthDate---|----Location----|---Phone--|---Email---|---DischargeDate---|--CustomerGroup---|
		*/
		if index == 0 {
			continue
		} else if index == index_data {
			return
		}
		var client Cliente = Cliente{}
		fmt.Println(row)
	}
}
