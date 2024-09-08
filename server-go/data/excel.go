package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"cristianvega6150/server/database"
	glob "cristianvega6150/server/globals"

	"github.com/xuri/excelize/v2"
)

/*
Está función es para poder sacar información dentro del excel, con estos datos luego
*/

func GetDataExcel() {

	var data []database.Clientes
	currentDir := glob.BASE_PROJECT
	file, errFileExcel := excelize.OpenFile(filepath.Join(currentDir, "\\data\\clientes.xlsx"))
	filename := filepath.Join("clients.json")
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return
	}

	if errFileExcel != nil {
		log.Fatal("Error al Abria la hoja del excel")
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
		client := database.Clientes{
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

	currentDir := glob.BASE_PROJECT
	filename := filepath.Join(currentDir, "\\data\\clients.json")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
		os.WriteFile(filename, jsondata, 0666)
	}

}

func GetDataJsonClients() []database.Clientes {

	currentDir := glob.BASE_PROJECT
	filename := filepath.Join(currentDir, "\\data\\clients.json")
	fmt.Println(filename)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		panic("No esta el archivo para recabar información")
	} else {
		data, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal("Hubo un error al cargar los datos")
		}
		var arrayClients []database.Clientes
		json.Unmarshal(data, &arrayClients)

		return arrayClients
	}
}
