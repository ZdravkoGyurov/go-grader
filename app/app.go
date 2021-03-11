package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"grader/app/config"
	"grader/executor"

	"go.mongodb.org/mongo-driver/mongo"
)

// https://medium.com/honestbee-tw-engineer/gracefully-shutdown-in-go-http-server-5f5e6b83da5a
// https://medium.com/@pinkudebnath/graceful-shutdown-of-golang-servers-using-context-and-os-signals-cc1fa2c55e97

// Application ...
type Application struct {
	cfg          config.Config
	stopExecutor executor.StopFunc
	server       *http.Server
	done         chan os.Signal
	dbClient     *mongo.Client
}

// New ...
func New(cfg config.Config, stopExecutor executor.StopFunc, handler http.Handler, dbClient *mongo.Client) *Application {
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return &Application{
		cfg:          cfg,
		stopExecutor: stopExecutor,
		server:       server,
		done:         done,
		dbClient:     dbClient,
	}
}

// Start ...
func (s *Application) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed listen and serve: %s\n", err)
		}
	}()
	log.Println("application started...")

	<-s.done

	defer func() {
		s.stopExecutor()
		log.Println("stopped job executor")

		ctx, cancel := context.WithTimeout(context.Background(), s.cfg.DBDisconnectTimeout)
		defer cancel()
		if err := s.dbClient.Disconnect(ctx); err != nil {
			log.Fatalf("failed to disconnect from mongodb: %s", err)
		}
		log.Println("disconnected from mongodb")
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown http server: %s\n", err)
	}
	log.Println("application stopped")
}

// Stop ...
func (s *Application) Stop() {
	s.done <- os.Interrupt
}
