package handler

import (
	"encoding/json"
	"net/http"

	"github.com/VictorArruda22/api-golang/internal/dto"
	"github.com/VictorArruda22/api-golang/internal/service"
)

type CategoryHandler struct {
	sv service.CategoryService
}

func NewCategoryHandler(sv service.CategoryService) CategoryHandler {
	return CategoryHandler{sv: sv}
}

func (c *CategoryHandler) CreateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCategory dto.CategoryRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
			http.Error(w, "Erro ao deserializar o JSON", http.StatusBadRequest)
			return
		}
		categoryResponseDTO, err := c.sv.CreateCategory(newCategory)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		responseJSON, err := json.Marshal(categoryResponseDTO)
		if err != nil {
			http.Error(w, "Erro ao processar a resposta", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Categoria criada com sucesso"))
		w.Write(responseJSON)
	}
}
