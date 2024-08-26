package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"cristianvega6150/server"

	"github.com/xuri/excelize/v2"
)

// Struct para representar a los clientes
type Cliente struct {
	ID                 string `json:"customer_id"`
	FullName           string `json:"full_name"`
	BirthDate          string `json:"birth_date"`
	CustomerAddress    string `json:"customer_address"`
	CustomerSector     string `json:"customer_sector"`
	CustomerPostalCode string `json:"customer_costal_code"`
	Phone              int    `json:"phone"`
	Email              string `json:"email"`
	DischargeDate      string `json:"discharge_date"`
	CustomerGroup      string `json:"customer_group"`
}

/*
Está función es para poder sacar información dentro del excel, con estos datos luego
*/

func GetDataExcel() {

	var data []Cliente
	currentDir, errPath := os.Getwd()
	file, errFileExcel := excelize.OpenFile(filepath.Join(currentDir, "\\data\\clientes.xlsx"))
	filename := filepath.Join(currentDir, "\\data\\clients.json")
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return
	}

	if errFileExcel != nil {
		log.Fatal("Error al Abria la hoja del excel")
	}
	if errPath != nil {
		log.Fatal("error en las rutas")
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, _ := file.GetRows("Hoja1")

	//Parsear la columna

	rowsParse := rows[1 : len(rows)-1]
	for _, row := range rowsParse {

		/*
			|---ID---|---FullName---|---BirthDate---|----Location----|---Phone--|---Email---|---DischargeDate---|--CustomerGroup---|
		*/

		sector_codePosta := strings.Split(row[4], ",")

		var sector string
		var postcode string
		if len(sector_codePosta) == 1 {
			sector = sector_codePosta[0]
			postcode = ""
		} else if len(sector_codePosta) == 2 {
			sector = sector_codePosta[0]
			postcode = sector_codePosta[1]
		}

		intphone, _ := strconv.Atoi(row[5])
		client := Cliente{
			ID:                 row[0],
			FullName:           row[1],
			BirthDate:          row[2],
			CustomerAddress:    row[3],
			CustomerSector:     sector,
			CustomerPostalCode: postcode,
			Phone:              intphone,
			Email:              row[6],
			DischargeDate:      row[7],
			CustomerGroup:      row[8],
		}

		new_data := append(data, client)
		data = new_data
	}
	json, errJson := json.Marshal(data)
	ClientsJson(json)

	if errJson != nil {
		log.Fatal("Not translate JSON")
	}

}

func ClientsJson(jsondata []byte) {
	currentDir, errPath := os.Getwd()
	if errPath != nil {
		log.Fatal("Error al mapear el archivo raiz")
	}

	filename := filepath.Join(currentDir, "\\data\\clients.json")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
		os.WriteFile(filename, jsondata, 0666)
	}

}

func GetDataJsonClients() []Cliente {
	currentDir, errPath := os.Getwd()
	if errPath != nil {
		log.Fatal("Error al mapear el archivo raiz")
	}

	filename := filepath.Join(currentDir, "\\data\\clients.json")

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		panic("No esta el archivo para recabar información")
	} else {
		data, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal("Hubo un error al cargar los datos")
			s
		}
		var arrayClients []Cliente
		json.Unmarshal(data, &arrayClients)

		return arrayClients
	}
}
