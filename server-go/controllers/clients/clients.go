package data

import (
	"cristianvega6150/server/database"
	errors_client "cristianvega6150/server/errors"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func DataController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	models := database.Cliente{}

	clientes, errorClientes := models.GetAllClients()

	if errorClientes != nil {
		log.Fatal(errorClientes.Error())
	}

	// pathvalue, err := strconv.Atoi(r.PathValue("total_index"))
	// fmt.Println(pathvalue)
	// if err != nil {
	// 	message_error_json, _ := json.Marshal(map[string]string{
	// 		"error": "No se permite digitos",
	// 	})

	// 	errors_client.ErrorClientJson(w, message_error_json, http.StatusBadRequest)
	// }

	jsonData, errJson := json.Marshal(clientes)
	if errJson != nil {
		message_error_json, _ := json.Marshal(map[string]string{
			"error": "No set transformo en json",
		})

		errors_client.ErrorClientJson(w, message_error_json, http.StatusBadRequest)

	}

	w.Write(jsonData)
}

func MessageController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Envía el mensaje
	fmt.Fprintf(w, "Este es un mensaje desde el controlador en Go")
}
