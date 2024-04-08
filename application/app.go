package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	return &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}
}

func (a *App) Start(ctx context.Context) error {
	s := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis %v", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		err = s.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server %v", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		fmt.Println("\nERROR! closing...")
		return err
	case <-ctx.Done():
		fmt.Println("\nSIGING received: closing...")
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return s.Shutdown(timeout)
	}
}
