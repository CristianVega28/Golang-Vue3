package data

import (
	"fmt"
	"net/http"
	"strconv"
)

func DataController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pathvalue, err := strconv.Atoi(r.PathValue("total_index"))
	fmt.Println(pathvalue)
	if err != nil {

		http.Error(w, "Solo esta permitido los digitos", http.StatusBadRequest)
	}
	// w.Write(data.GetData(pathvalue))

}

func DataPostController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("metodo post")
}
