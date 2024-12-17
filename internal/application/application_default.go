package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitApplication(mux *chi.Mux) {
	err := http.ListenAndServe(":"+"8080", mux)
	if err != nil {
		panic(err)
	}


}