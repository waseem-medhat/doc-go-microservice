package application

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wipdev-tech/doc-go-microservice/handler"
)

func loadRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/orders", loadOrderRoutes)

	return r
}

func loadOrderRoutes(router chi.Router) {
	h := &handler.Order{}

	router.Post("/", h.Create)
	router.Get("/", h.List)
	router.Get("/{id}", h.GetByID)
	router.Put("/{id}", h.UpdateByID)
	router.Delete("/{id}", h.DeleteByID)
}
