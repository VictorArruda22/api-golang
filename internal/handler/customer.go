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

type CustomerHandler struct {
	sv *service.CustomerService
}

func NewCustomerHandler(sv *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{sv: sv}
}

func (h *CustomerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := h.sv.GetAll()
		if err != nil {
			code := http.StatusInternalServerError
			msg := utils.ErrCostumerRepositoryInternalError.Error()
			utils.ResponseJSON(w, code, msg, nil)
			return
		}
		utils.ResponseJSON(w, http.StatusOK, "Customers encontrados", customers)
	}
}

func (h *CustomerHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrCostumerRepositoryInvalidID.Error(), nil)
			return
		}

		customer, err := h.sv.GetByID(id)
		if err != nil {
			code := http.StatusBadRequest
			msg := err.Error()
			if errors.Is(err, utils.ErrCostumerRepositoryNotFound) {
				code = http.StatusNotFound
			}
			utils.ResponseJSON(w, code, msg, nil)
			return
		}
		utils.ResponseJSON(w, http.StatusOK, "Success", customer)
	}
}

func (h *CustomerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer entities.Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrCostumerRepositoryRequest.Error(), nil)
			return
		}

		createdCustomer, err := h.sv.Create(customer)
		if err != nil {
			code := http.StatusBadRequest
			msg := err.Error()
			if errors.Is(err, utils.ErrCostumerRepositoryBadField) {
				code = http.StatusBadRequest
			} else if errors.Is(err, utils.ErrCostumerRepositoryNullValue) {
				code = http.StatusUnprocessableEntity
			}
			utils.ResponseJSON(w, code, msg, nil)
			return
		}
		utils.ResponseJSON(w, http.StatusOK, "Customer Created", createdCustomer)
	}
}

func (h *CustomerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrCostumerRepositoryInvalidID.Error(), nil)
			return
		}

		var customer entities.Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrCostumerRepositoryRequest.Error(), nil)
			return
		}

		customer.Id = id

		if _, err := h.sv.Update(customer); err != nil {
			code := http.StatusBadRequest
			msg := err.Error()
			if errors.Is(err, utils.ErrCostumerRepositoryNotFound) {
				code = http.StatusNotFound
			} else if errors.Is(err, utils.ErrCostumerRepositoryBadField) {
				code = http.StatusBadRequest
			} else if errors.Is(err, utils.ErrCostumerRepositoryNullValue) {
				code = http.StatusUnprocessableEntity
			}
			utils.ResponseJSON(w, code, msg, nil)
			return
		}
		utils.ResponseJSON(w, http.StatusOK, "Customer updated", customer)
	}
}

func (h *CustomerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrCostumerRepositoryInvalidID.Error(), nil)
			return
		}

		if err := h.sv.Delete(id); err != nil {
			code := http.StatusBadRequest
			msg := err.Error()
			if errors.Is(err, utils.ErrCostumerRepositoryNotFound) {
				code = http.StatusNotFound
			}
			utils.ResponseJSON(w, code, msg, nil)
			return
		}
		utils.ResponseJSON(w, http.StatusNoContent, "Customer deleted", nil)
	}
}
