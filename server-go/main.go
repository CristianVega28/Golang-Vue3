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

	//TODO
	//* Segundo argumento -> Limpiar toda la tabla de clientes
	//* Tercer argumento -> Colocar cuantos registros inyectar a la tabla clientes

	if len(args) > 1 {
		if args[1] == "artisan" {

			services := database.ServiceDB{}
			db, err := services.Conection()
			if args[2] == "--seed" {
				if err != nil {
					panic("error a la conexion de base de datos")
				}
				data := excel.GetDataJsonClients()
				seeders.ClientSeedStart(data, db)

			}
			if args[2] == "--clear" {
				if err != nil {
					panic("error a la conexion de base de datos")
				}
				clients := database.Clientes{}
				message, errDelete := clients.DeleteAllClients()
				if errDelete != nil {
					panic(errDelete.Error())
				}
				fmt.Println(*message)
			}

			fmt.Println("No existe ese comando")

			return
		}

	}

	server_core := core.Server{
		Config: core.ConfigServerMux{
			Address: ":3000",
		},
		Handler: map[string]func(http.ResponseWriter, *http.Request){
			"GET /clients":  data.DataController,
			"POST /clients": data.ByIdCliente,
			"GET /":         data.MessageController,
			"POST /form":    forms.FormContrller,
		},
	}
	excel.GetDataExcel()
	server_core.Start()

}
