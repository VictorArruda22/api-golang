package handler

import (
	"encoding/json"
	"errors"
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
		authors, err := h.Service.GetAll()
		if err != nil {
			code := http.StatusInternalServerError
			msg := utils.ErrAuthorRepositoryInternalError.Error()
			utils.ResponseJSON(w, code, msg, nil)
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
			code := http.StatusBadRequest
			msg := err.Error()
			if errors.Is(err, utils.ErrAuthorRepositoryNotFound) {
				code = http.StatusNotFound
			}
			utils.ResponseJSON(w, code, msg, nil)
			return
		}
		utils.ResponseJSON(w, http.StatusOK, "Autor encontrado", author)
	}
}

func (h *AuthorHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author entities.Author
		if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrAuthorRepositoryRequest.Error(), nil)
			return
		}

		createdAuthor, err := h.Service.Create(author)
		if err != nil {
			code := http.StatusBadRequest
			msg := err.Error()
			if errors.Is(err, utils.ErrAuthorRepositoryBadField) {
				code = http.StatusBadRequest
			} else if errors.Is(err, utils.ErrAuthorRepositoryNullValue) {
				code = http.StatusBadRequest
			}
			utils.ResponseJSON(w, code, msg, nil)
			return
		}
		utils.ResponseJSON(w, http.StatusOK, "Autor criado.", map[string]int{"id": createdAuthor[0].ID})
	}
}

func (h *AuthorHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrAuthorRepositoryInvalidID.Error(), nil)
			return
		}
		var author entities.Author
		if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrAuthorRepositoryRequest.Error(), nil)
			return
		}

		author.ID = id

		if _, err := h.Service.Update(author); err != nil {
			code := http.StatusBadRequest
			msg := err.Error()
			if errors.Is(err, utils.ErrAuthorRepositoryRequest) {
				code = http.StatusBadRequest
			} else if errors.Is(err, utils.ErrAuthorRepositoryBadField) {
				code = http.StatusBadRequest
			} else if errors.Is(err, utils.ErrAuthorRepositoryNotFound) {
				code = http.StatusNotFound
			}

			utils.ResponseJSON(w, code, msg, nil)
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
			code := http.StatusBadRequest
			msg := err.Error()
			if errors.Is(err, utils.ErrAuthorRepositoryNotFound) {
				code = http.StatusNotFound
			}
			utils.ResponseJSON(w, code, msg, nil)
			return
		}

		utils.ResponseJSON(w, http.StatusNoContent, "Autor deletado com sucesso.", nil)
	}
}
