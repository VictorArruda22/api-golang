package routers

import (
	"github.com/VictorArruda22/api-golang/internal/handler"
	"github.com/go-chi/chi/v5"
)

func RouterCustomerManagement(customerHandler handler.CustomerHandler) *chi.Mux {

	router := chi.NewRouter()

	router.Get("/customers", customerHandler.GetAll())
	router.Get("/customers/{id}", customerHandler.GetByID())
	router.Post("/customers", customerHandler.Create())
	router.Put("/customers/{id}", customerHandler.Update())
	router.Delete("/customers/{id}", customerHandler.Delete())

	return router

}
