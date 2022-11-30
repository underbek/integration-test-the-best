package http_router

import (
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/underbek/integration-test-the-best/user-service/internal/transport/http-handler"
)

func New(h *handler.Handler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/users", h.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", h.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users/deposit", h.DepositBalance).Methods(http.MethodPost)

	return r
}
