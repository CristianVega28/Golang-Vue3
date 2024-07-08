package main

import (
	"cristianvega6150/server/controllers/data"
	"cristianvega6150/server/controllers/forms"
	"cristianvega6150/server/core"
	json "cristianvega6150/server/data"
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	server_core := core.Server{
		Config: core.ConfigServerMux{
			Address: ":3000",
		},
		Handler: map[string]func(http.ResponseWriter, *http.Request){
			"GET /data":  data.DataController,
			"POST /form": forms.FormContrller,
		},
	}
	fmt.Println(filepath.Clean("/clientes.xlsx"))
	json.GetData()
	server_core.Start()
}
