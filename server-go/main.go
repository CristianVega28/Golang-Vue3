package main

import (
	data "cristianvega6150/server/controllers/clients"
	"cristianvega6150/server/controllers/forms"
	"cristianvega6150/server/core"
	excel "cristianvega6150/server/data"
	"cristianvega6150/server/database"
	"cristianvega6150/server/database/seeders"
	"net/http"
)

func main() {
	server_core := core.Server{
		Config: core.ConfigServerMux{
			Address: ":3000",
		},
		Handler: map[string]func(http.ResponseWriter, *http.Request){
			"GET /data/{total_index}": data.DataController,
			"GET /":                   data.MessageController,
			"POST /form":              forms.FormContrller,
		},
	}
	excel.GetDataExcel()
	services := database.ServiceDB{}
	db, err := services.Conection()
	if err != nil {
		panic("error a la conexion de base de datos")
	}
	data := excel.GetDataJsonClients()
	// fmt.Println(data)
	seeders.ClientSeedStart(data, db)
	server_core.Start()

}
