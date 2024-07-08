package forms

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func FormContrller(w http.ResponseWriter, r *http.Request) {
	var person Person

	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		http.Error(w, "Error en la decodificación del usuario", http.StatusBadRequest)
		return
	}

	fmt.Printf("Nuevo usuario %+v", person)
}
