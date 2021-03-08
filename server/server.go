package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// https://medium.com/honestbee-tw-engineer/gracefully-shutdown-in-go-http-server-5f5e6b83da5a
// https://medium.com/@pinkudebnath/graceful-shutdown-of-golang-servers-using-context-and-os-signals-cc1fa2c55e97

// Server ...
type Server struct {
	server *http.Server
	done   chan os.Signal
	client *mongo.Client
}

// New ...
func New(address string, handler http.Handler, client *mongo.Client) *Server {
	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return &Server{
		server: server,
		done:   done,
		client: client,
	}
}

// Start ...
func (s *Server) Start() {
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
		if err := s.client.Disconnect(ctx); err != nil {
			log.Fatalf("failed to disconnect from mongodb: %s", err)
		}
		log.Println("disconnected from mongodb...")

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
