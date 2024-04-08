package main

import (
	"context"
	"fmt"

	"github.com/wipdev-tech/doc-go-microservice/application"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
}
