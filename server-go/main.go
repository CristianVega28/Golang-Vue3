package main

import (
	"cristianvega6150/server/controllers/data"
	"cristianvega6150/server/core"
	"net/http"
)

func main() {
	server_core := core.Server{
		Config: core.ConfigServerMux{
			Address: ":3000",
		},
		Handler: map[string]func(http.ResponseWriter, *http.Request){
			"GET /data":  data.DataController,
			"POST /data": data.DataPostController,
		},
	}

	server_core.Start()
}
