package main

import (
	"log"
	"net/http"

	"github.com/underbek/integration-test-the-best/user_service/migrate"

	"github.com/underbek/integration-test-the-best/user_service/billing"
	"github.com/underbek/integration-test-the-best/user_service/handler"
	"github.com/underbek/integration-test-the-best/user_service/server"
	"github.com/underbek/integration-test-the-best/user_service/storage"
	"github.com/underbek/integration-test-the-best/user_service/use_case"
)

const (
	addr        = ":8080"
	dbDsn       = "host=localhost port=5432 user=user password=password dbname=postgres sslmode=disable"
	billingAddr = "http://localhost:8085"
)

func main() {
	repo, err := storage.New(dbDsn)
	if err != nil {
		log.Fatal(err)
	}
	err = migrate.Migrate(dbDsn, migrate.Migrations)
	if err != nil {
		log.Fatal(err)
	}
	billingClient := billing.New(http.DefaultClient, billingAddr)
	useCase := use_case.New(repo, billingClient)
	h := handler.New(useCase)
	srv := server.New(addr, h)
	log.Fatal(srv.Serve())
}
