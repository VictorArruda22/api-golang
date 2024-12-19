package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/VictorArruda22/api-golang/internal/entities"
	"github.com/VictorArruda22/api-golang/internal/service"
	"github.com/VictorArruda22/api-golang/internal/utils"
	"github.com/go-chi/chi/v5"
)

type AuthorHandler struct {
	Service *service.AuthorService
}

func NewAuthorHandler(service *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{Service: service}
}

func (h *AuthorHandler) GetAllAuthors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handler: GetAll called")
		authors, err := h.Service.GetAll()
		if err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, utils.ErrAuthorRepositoryInternalError.Error(), nil)
			return
		}
		utils.ResponseJSON(w, http.StatusOK, "Autores encontrados", authors)
	}
}

func (h *AuthorHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrAuthorRepositoryInvalidID.Error(), nil)
			return
		}

		author, err := h.Service.GetByID(id)
		if err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, utils.ErrAuthorRepositoryInternalError.Error(), nil)
			return
		}
		utils.ResponseJSON(w, http.StatusOK, "Autor encontrado", author)
	}
}

func (h *AuthorHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author entities.Author
		if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrAuthorrRepositoryRequest.Error(), nil)
			return
		}

		createdAuthor, err := h.Service.Create(author)
		if err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, utils.ErrAuthorRepositoryInternalError.Error(), nil)
			return
		}
		utils.ResponseJSON(w, http.StatusOK, "Autor criado.", map[string]int{"id": createdAuthor[0].ID})
	}
}

func (h *AuthorHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author entities.Author
		if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrAuthorrRepositoryRequest.Error(), nil)
			return
		}
		if _, err := h.Service.Update(author); err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, utils.ErrAuthorRepositoryInternalError.Error(), nil)
			return
		}
		utils.ResponseJSON(w, http.StatusOK, "Autor atualizado.", nil)
	}
}

func (h *AuthorHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrAuthorRepositoryInvalidID.Error(), nil)
			return
		}

		if err := h.Service.Delete(id); err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, utils.ErrAuthorRepositoryInternalError.Error(), nil)
			return
		}

		utils.ResponseJSON(w, http.StatusNoContent, "Autor deletado com sucesso.", nil)
	}
}
