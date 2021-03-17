package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ZdravkoGyurov/go-grader/pkg/app/config"
	"github.com/ZdravkoGyurov/go-grader/pkg/executor"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/storage"
)

// Context ...
type Context struct {
	Context context.Context
	Cancel  context.CancelFunc
	Cfg     config.Config
}

// NewContext ...
func NewContext() Context {
	cfg := config.Config{
		Host:                      "0.0.0.0",
		Port:                      8080,
		ServerReadTimeout:         30 * time.Second,
		ServerWriteTimeout:        30 * time.Second,
		MaxExecutorWorkers:        5,
		MaxExecutorConcurrentJobs: 100,
		DatabaseURI:               "mongodb://host.docker.internal:27017",
		DBConnectTimeout:          30 * time.Second,
		DBDisconnectTimeout:       30 * time.Second,
		DatabaseName:              "grader",
		ServerShutdownTimeout:     5 * time.Second,
		SessionCookieName:         "Grader",
		TestsGitUser:              "ZdravkoGyurov",
		TestsGitRepo:              "grader-docker-tests",
	}

	ctx, cancel := context.WithCancel(context.Background())

	return Context{
		Context: ctx,
		Cancel:  cancel,
		Cfg:     cfg,
	}
}

// Application ...
type Application struct {
	appContext Context
	exec       *executor.Executor
	storage    *storage.Storage
	server     *http.Server
}

// New ...
func New(appContext Context, exec *executor.Executor, storage *storage.Storage, handler http.Handler) *Application {
	address := fmt.Sprintf("%s:%d", appContext.Cfg.Host, appContext.Cfg.Port)
	server := &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  appContext.Cfg.ServerReadTimeout,
		WriteTimeout: appContext.Cfg.ServerWriteTimeout,
	}

	return &Application{
		appContext: appContext,
		exec:       exec,
		storage:    storage,
		server:     server,
	}
}

// Start ...
func (a *Application) Start() {
	a.setupSignalNotifier()

	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Fatalf("failed listen and serve: %s\n", err)
		}
	}()
	log.Info().Println("Application started...")

	<-a.appContext.Context.Done()

	a.stopExecutor()
	a.disconnectFromDB()
	a.shutdownServer()
	log.Info().Println("Application stopped gracefully")
}

// Stop ...
func (a *Application) Stop() {
	a.appContext.Cancel()
}

func (a *Application) setupSignalNotifier() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		log.Info().Println("Stopping application...")
		a.appContext.Cancel()
	}()
}

func (a *Application) stopExecutor() {
	a.exec.Stop()
	log.Info().Println("Executor stopped")
}

func (a *Application) disconnectFromDB() {
	a.storage.Disconnect()
	log.Info().Println("Disconnected from DB")
}

func (a *Application) shutdownServer() {
	ctx, cancel := context.WithTimeout(context.Background(), a.appContext.Cfg.ServerShutdownTimeout)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		log.Error().Fatalf("failed to shutdown http server: %s\n", err)
	}
	log.Info().Println("Server shutdown")
}
