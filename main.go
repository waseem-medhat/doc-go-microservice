package main

import (
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Get("/hello", basicHandler)
	router.Use(middleware.Logger)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hiyaa!!"))
}
