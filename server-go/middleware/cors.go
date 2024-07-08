package middleware

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Si es una solicitud OPTIONS, solo respondemos con los headers CORS permitidos y no continuamos con el siguiente handler
		if r.Method == "OPTIONS" {
			return
		}

		// Llamada al siguiente handler
		next.ServeHTTP(w, r)
	})
}
