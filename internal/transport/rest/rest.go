package rest

import (
	"context"
	"eth-parser/internal/config"
	"eth-parser/internal/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const maxHeaderBytes = 1 << 20

type Server struct {
	srv      *http.Server
	services *service.Service
}

func New(services *service.Service) *Server {
	return &Server{
		services: services,
	}
}

func (s *Server) Run(ctx context.Context, cfg *config.Config, handler http.Handler) error {
	s.srv = &http.Server{
		Addr:           ":" + cfg.HTTP.Port,
		Handler:        handler,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: maxHeaderBytes,
	}

	if err := s.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to run the HTTP server: %w", err)
	}

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := s.srv.Shutdown(shutdownCtx); err != nil {
			return
		}
	}()

	return nil
}

func (s *Server) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/addresses", s.addAddress).Methods("POST")

	return r
}
