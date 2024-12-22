package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/VictorArruda22/api-golang/internal/dto"
	"github.com/VictorArruda22/api-golang/internal/entities"
	"github.com/VictorArruda22/api-golang/internal/utils"
	"github.com/go-chi/chi/v5"
)

type CategoryHandler struct {
	sv entities.CategoryService
}

func NewCategoryHandler(sv entities.CategoryService) *CategoryHandler {
	return &CategoryHandler{sv: sv}
}

func (c *CategoryHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCategory dto.CategoryRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
			utils.ErrorResponse(w, utils.ErrCategoryRepositoryRequest.Error(), http.StatusUnprocessableEntity)
			return
		}
		categoryResponseDTO, err := c.sv.Create(newCategory)
		if err != nil {
			utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, http.StatusCreated, "Categoria criada com sucesso", categoryResponseDTO)
	}
}

func (c *CategoryHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ErrorResponse(w, utils.ErrCategoryRepositoryInvalidID.Error(), http.StatusBadRequest)
			return
		}
		err = c.sv.Delete(id)
		if err != nil && errors.Is(err, utils.ErrCategoryRepositoryNotFound) {
			utils.ErrorResponse(w, err.Error(), http.StatusNotFound)
			return
		}
		if err != nil && errors.Is(err, utils.ErrCategoryRepositoryInternalError) {
			utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, http.StatusOK, "Categoria deletada com sucesso", nil)
	}
}

func (c *CategoryHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ErrorResponse(w, utils.ErrCategoryRepositoryInvalidID.Error(), http.StatusBadRequest)
			return
		}

		categoryResponseDTO, err := c.sv.GetByID(id)
		if err != nil && errors.Is(err, utils.ErrCategoryRepositoryNotFound) {
			utils.ErrorResponse(w, err.Error(), http.StatusNotFound)
			return
		}
		if err != nil && errors.Is(err, utils.ErrCategoryRepositoryInternalError) {
			utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, http.StatusOK, "Categoria selecionada", categoryResponseDTO)
	}
}

func (c *CategoryHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var categories []dto.CategoryResponseDTO
		categories, err := c.sv.GetAll()
		if err != nil {
			utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, http.StatusOK, "Categorias selecionadas", categories)
	}
}

func (c *CategoryHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ErrorResponse(w, utils.ErrCategoryRepositoryInvalidID.Error(), http.StatusBadRequest)
			return
		}
		var categoryUpdated dto.CategoryRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&categoryUpdated); err != nil {
			utils.ErrorResponse(w, utils.ErrCategoryRepositoryRequest.Error(), http.StatusUnprocessableEntity)
			return
		}

		categoryUpdatedResponseDTO, err := c.sv.Update(id, categoryUpdated)
		if err != nil && errors.Is(err, utils.ErrCategoryRepositoryNotFound) {
			utils.ErrorResponse(w, err.Error(), http.StatusNotFound)
			return
		}
		if err != nil && errors.Is(err, utils.ErrCategoryRepositoryInternalError) {
			utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, http.StatusOK, "Categoria atualizada com sucesso", categoryUpdatedResponseDTO)
	}
}
