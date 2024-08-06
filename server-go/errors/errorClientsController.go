package errors

import "net/http"

func ErrorClientJson(w http.ResponseWriter, errors []byte, status int) {
	writeJsonError(w, status)

	w.Write(errors)
}

func writeJsonError(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}
