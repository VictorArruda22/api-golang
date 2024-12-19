package routers

import (
	"fmt"
	"net/http"

	"github.com/VictorArruda22/api-golang/internal/handler"
	"github.com/go-chi/chi/v5"
)

func RouterManagement(categoryHandler handler.CategoryHandler) *chi.Mux {

	router := chi.NewRouter()

	router.Get("/", (func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}))

	router.Post("/category", categoryHandler.CreateCategory())

	return router

}
