package data

import (
	"fmt"
	"net/http"
)

func DataController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryUrl := r.URL.Query()

	service := ServiceCliente{}
	if queryUrl.Has("page") && queryUrl.Has("amount") {
		service.paginationService(w, r)
	} else {
		service.allDataService(w, r)
	}

}

func ByIdCliente(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	service := ServiceCliente{}
	service.getById(w, r)
}

func MessageController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Envía el mensaje
	fmt.Fprintf(w, "Este es un mensaje desde el controlador en Go")
}
