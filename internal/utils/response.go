package utils

import (
	"encoding/json"
	"net/http"
)

// Response padrão para a API
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Função utilitária para enviar respostas JSON
func ResponseJSON(w http.ResponseWriter, status int, message string, data interface{}) {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
