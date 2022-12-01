package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/underbek/integration-test-the-best/user-service/internal/config"
	"github.com/underbek/integration-test-the-best/user-service/internal/repository/billing"
	"github.com/underbek/integration-test-the-best/user-service/internal/repository/storage"
	httpHandler "github.com/underbek/integration-test-the-best/user-service/internal/transport/http-handler"
	httpRouter "github.com/underbek/integration-test-the-best/user-service/internal/transport/http-router"
	useCase "github.com/underbek/integration-test-the-best/user-service/internal/use-case"
	"github.com/underbek/integration-test-the-best/user-service/migrate"
)

type App struct {
	HttpServer        *http.Server
	BillingHttpClient *http.Client
}

func New(cfg config.Config) (*App, error) {

	// запускаем миграции
	err := migrate.Migrate(cfg.PostgresDSN, migrate.Migrations)
	if err != nil {
		return nil, err
	}

	// создаем сторадж
	repo, err := storage.New(cfg.PostgresDSN)
	if err != nil {
		return nil, err
	}

	billingHttpClient := http.DefaultClient

	// создаем клиент к биллинг сервису
	billingClient := billing.New(billingHttpClient, cfg.BillingDSN)

	// создаем use-case
	uc := useCase.New(repo, billingClient)

	// создаем handler
	h := httpHandler.New(uc)

	// создаем router
	r := httpRouter.New(h)

	// создаем server
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}

	return &App{
		HttpServer:        srv,
		BillingHttpClient: billingHttpClient,
	}, nil
}

// Run запускает все приложение
func (a *App) Run() error {
	log.Println("starting service")
	return a.HttpServer.ListenAndServe()
}
