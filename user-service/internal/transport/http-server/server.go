package http_server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/underbek/integration-test-the-best/user-service/handler"
)

type Server struct {
	server *http.Server
	Router *mux.Router
}

func New(addr string, h *handler.Handler) *Server {
	r := mux.NewRouter()
	r.HandleFunc("/users", h.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", h.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users/deposit", h.DepositBalance).Methods(http.MethodPost)

	server := &http.Server{
		Handler: r,
		Addr:    addr,
	}

	return &Server{
		server: server,
		Router: r,
	}
}

func (s *Server) Serve() error {
	return s.server.ListenAndServe()
}
