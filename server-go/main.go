package main

import (
	data "cristianvega6150/server/controllers/clients"
	"cristianvega6150/server/controllers/forms"
	"cristianvega6150/server/core"
	excel "cristianvega6150/server/data"
	"net/http"
)

func main() {
	server_core := core.Server{
		Config: core.ConfigServerMux{
			Address: ":3000",
		},
		Handler: map[string]func(http.ResponseWriter, *http.Request){
			"GET /data/{total_index}": data.DataController,
			"POST /form":              forms.FormContrller,
		},
	}
	excel.GetDataExcel()
	excel.GetDataJsonClients()
	server_core.Start()

}
