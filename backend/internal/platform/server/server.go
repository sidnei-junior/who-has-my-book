// Package server wires HTTP routes and middlewares and runs the HTTP server.
package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"gorm.io/gorm"

	"github.com/sidnei-junior/who-has-my-book/backend/internal/platform/database"
)

// Server is the API HTTP server.
type Server struct {
	httpServer *http.Server
	db         *gorm.DB
	log        *slog.Logger
}

// New builds the server with all routes registered.
func New(port int, allowedOrigins []string, db *gorm.DB, log *slog.Logger) *Server {
	s := &Server{db: db, log: log}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.handleHealth)

	var handler http.Handler = mux
	handler = cors(allowedOrigins)(handler)
	handler = requestLogger(log)(handler)
	handler = recoverer(log)(handler)

	s.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	return s
}

// Start runs the server until it is shut down. It returns nil on graceful shutdown.
func (s *Server) Start() error {
	s.log.Info("http server listening", "addr", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Shutdown gracefully stops the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := database.Ping(ctx, s.db); err != nil {
		s.log.Error("health check failed", "error", err)
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "unhealthy", "database": "down"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "database": "up"})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
