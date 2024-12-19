package routers

import (
	"github.com/VictorArruda22/api-golang/internal/handler"
	"github.com/go-chi/chi/v5"
)

func RouterManagement(authorHandler handler.AuthorHandler) *chi.Mux {

	router := chi.NewRouter()

	router.Get("/authors", authorHandler.GetAllAuthors())
	router.Get("/authors/{id}", authorHandler.GetByID())
	router.Post("/authors", authorHandler.Create())
	router.Put("/authors/{id}", authorHandler.Update())
	router.Delete("/authors/{id}", authorHandler.Delete())

	return router

}
