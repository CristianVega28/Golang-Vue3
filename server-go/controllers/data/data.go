package data

import (
	"fmt"
	"net/http"
)

func DataController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Esto es una prueba"))
}

func DataPostController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("metodo post")
}
