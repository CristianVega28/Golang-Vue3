package data

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"cristianvega6150/server/data"
)

func DataController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	pathvalue, err := strconv.Atoi(r.PathValue("total_index"))

	if err != nil {
		errors.New("El indice no es un número")
	}
	w.Write(data.GetData(pathvalue))
}

func DataPostController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("metodo post")
}
