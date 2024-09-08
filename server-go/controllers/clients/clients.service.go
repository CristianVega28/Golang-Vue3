package data

import (
	"cristianvega6150/server/database"
	errors_client "cristianvega6150/server/errors"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ServiceCliente struct {
}

func (ctx *ServiceCliente) allDataService(w http.ResponseWriter, r *http.Request) {
	models := database.Clientes{}

	clientes, errorClientes := models.GetAllClients()

	if errorClientes != nil {
		log.Fatal(errorClientes.Error())
	}

	jsonData, errJson := json.Marshal(clientes)
	if errJson != nil {
		message_error_json, _ := json.Marshal(map[string]string{
			"error": "No set transformo en json",
		})

		errors_client.ErrorClientJson(w, message_error_json, http.StatusBadRequest)

	}

	w.Write(jsonData)
}

func (ctx *ServiceCliente) paginationService(w http.ResponseWriter, r *http.Request) {

	queryUrl := r.URL.Query()
	page, errPage := strconv.Atoi(queryUrl.Get("page"))
	amount, errAmount := strconv.Atoi(queryUrl.Get("amount"))

	if errPage != nil || errAmount != nil {
		errors.New("no se pudo convertir")
	}

	models := database.Clientes{}

	dataPagination, count, err := models.Pagination(page, amount)

	if err != nil {
		fmt.Println(err.Error())
	}

	responsePagination, errJson := json.Marshal(map[string]any{
		"data_pagination": dataPagination,
		"total_rows":      *count,
	})

	if errJson != nil {
		errors.New("error en parsear a json")
	}

	w.Write(responsePagination)
}
