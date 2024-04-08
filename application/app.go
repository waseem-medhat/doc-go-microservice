package application

import (
	"context"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	return &App{
		router: loadRoutes(),
	}
}

func (a *App) Start(ctx context.Context) error {
	s := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}
	return s.ListenAndServe()
}
