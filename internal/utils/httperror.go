package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse padrão para erros da API
func ErrorResponse(w http.ResponseWriter, message string, status int) {
	errorResponse := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse)
}
