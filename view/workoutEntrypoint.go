package view

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Route() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", helloworld)
	return r
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! 2"))
}

func generateWorkout(w http.ResponseWriter, r *http.Request) {
	r.
}