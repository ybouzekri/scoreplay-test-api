package rest

import (
	"log/slog"
	"net/http"
)

type Server struct {
	addr   string
	router http.Handler
	logger *slog.Logger
}

func NewServer(addr string, router http.Handler, logger *slog.Logger) *Server {
	return &Server{
		addr:   addr,
		router: router,
		logger: logger,
	}
}

func (s *Server) ListenAndServe() error {
	server := &http.Server{
		Addr:    s.addr,
		Handler: s.router,
	}

	s.logger.Info("start listening incoming connections", "addr", s.addr)

	return server.ListenAndServe()
}
