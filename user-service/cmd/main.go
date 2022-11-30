package main

import (
	"log"

	"github.com/underbek/integration-test-the-best/user-service/internal/app"
	"github.com/underbek/integration-test-the-best/user-service/internal/config"
)

func main() {
	cfg := config.New()

	app, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
