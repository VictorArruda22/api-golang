package routers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouterManagement() *chi.Mux {

	router := chi.NewRouter()

	router.Get("/", (func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}))

	return router

}