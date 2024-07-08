package core

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type RouterStruct struct {
}

// Configuración que se copia del http.Server
type ConfigServerMux struct {
	Address string
}

// Centralizando los valores
// Config -> Valores que colocar dentro de la inicialización del http.Server de la funcion Start
// Handler -> Asignación que manejará toda las rutas que se coloquen y sus respectivos controladores
type Server struct {
	Config  ConfigServerMux
	Handler map[string]func(http.ResponseWriter, *http.Request)
}

// Emepzará con el servidor una vez se llamé
func (server *Server) Start() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	routes := http.NewServeMux()

	for path, handler := range server.Handler {
		routes.HandleFunc(path, handler)
	}

	// handler := cors.Default().Handler(routes)
	http_server := &http.Server{
		Addr:    server.Config.Address,
		Handler: c.Handler(routes),
	}

	http_server.ListenAndServe()
	fmt.Println("liston to")
}
