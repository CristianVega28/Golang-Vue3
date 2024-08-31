package main

import (
	data "cristianvega6150/server/controllers/clients"
	"cristianvega6150/server/controllers/forms"
	"cristianvega6150/server/core"
	excel "cristianvega6150/server/data"
	"cristianvega6150/server/database"
	"cristianvega6150/server/database/seeders"
	"fmt"
	"net/http"
	"os"
)

func main() {

	args := os.Args

	if len(args) > 1 {
		if args[1] == "--seed" {
			services := database.ServiceDB{}
			db, err := services.Conection()
			if err != nil {
				panic("error a la conexion de base de datos")
			}
			data := excel.GetDataJsonClients()
			seeders.ClientSeedStart(data, db)

		} else {
			fmt.Println("No existe ese comando")
		}

	}

	server_core := core.Server{
		Config: core.ConfigServerMux{
			Address: ":3000",
		},
		Handler: map[string]func(http.ResponseWriter, *http.Request){
			"GET /clients": data.DataController,
			"GET /":        data.MessageController,
			"POST /form":   forms.FormContrller,
		},
	}
	excel.GetDataExcel()
	server_core.Start()

}
