package handler

import (
	"net/http"

	"github.com/VictorArruda22/api-golang/internal/service"
	"github.com/VictorArruda22/api-golang/internal/utils"
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
