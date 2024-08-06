package data

import (
	errors_client "cristianvega6150/server/errors"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func DataController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pathvalue, err := strconv.Atoi(r.PathValue("total_index"))
	fmt.Println(pathvalue)
	if err != nil {
		message_error_json, _ := json.Marshal(map[string]string{
			"error": "No se permite digitos",
		})

		errors_client.ErrorClientJson(w, message_error_json, http.StatusBadRequest)
	}
	// w.Write(data.GetData(pathvalue))
}

func DataPostController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("metodo post")
}
