package routers

import (
	"fmt"
	"net/http"

	"github.com/VictorArruda22/api-golang/internal/handler"
	"github.com/go-chi/chi/v5"
)

func RouterManagement(authorHandler handler.AuthorHandler) *chi.Mux {

	router := chi.NewRouter()

	router.Get("/", (func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}))

	router.Get("/authors", authorHandler.GetAllAuthors())
	router.Get("/authors/{id}", authorHandler.GetByID())
	router.Post("/authors", authorHandler.Create())
	router.Put("/authors/{id}", authorHandler.Update())
	router.Delete("/authors/{id}", authorHandler.Delete())

	return router

}
