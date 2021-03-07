package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	server *http.Server
	done   chan os.Signal
}

// New ...
func New(address string, handler http.Handler) *Server {
	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return &Server{
		server: server,
		done:   done,
	}
}

// Start ...
func (s *Server) Start(ctx context.Context) {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed listen and serve: %s\n", err)
		}
	}()
	log.Println("server started...")

	<-s.done
	log.Println("stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %s\n", err)
	}
	log.Println("server stopped")
}

// Stop ...
func (s *Server) Stop() {
	s.done <- os.Interrupt
}
