package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZdravkoGyurov/go-grader/app/config"
	"github.com/ZdravkoGyurov/go-grader/executor"
	"github.com/ZdravkoGyurov/go-grader/log"
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
		Host:                      "localhost",
		Port:                      8080,
		MaxExecutorWorkers:        5,
		MaxExecutorConcurrentJobs: 100,
		GithubTestsRepo:           "",
		DatabaseURI:               "mongodb://localhost:27017",
		DBConnectTimeout:          30 * time.Second,
		DBDisconnectTimeout:       30 * time.Second,
		DatabaseName:              "grader",
		ServerShutdownTimeout:     5 * time.Second,
		SessionCookieName:         "Grader",
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
	appCtx   Context
	exec     *executor.Executor
	dbClient *mongo.Client
	server   *http.Server
}

// New ...
func New(appCtx Context, exec *executor.Executor, dbClient *mongo.Client, handler http.Handler) *Application {
	address := fmt.Sprintf("%s:%d", appCtx.Cfg.Host, appCtx.Cfg.Port)
	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	return &Application{
		appCtx:   appCtx,
		exec:     exec,
		dbClient: dbClient,
		server:   server,
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

	<-a.appCtx.Context.Done()

	a.stopExecutor()
	a.disconnectFromDB()
	a.shutdownServer()
	log.Info().Println("Application stopped gracefully")
}

// Stop ...
func (a *Application) Stop() {
	a.appCtx.Cancel()
}

func (a *Application) setupSignalNotifier() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		log.Info().Println("Stopping application...")
		a.appCtx.Cancel()
	}()
}

func (a *Application) stopExecutor() {
	a.exec.Stop()
	log.Info().Println("Executor stopped")
}

func (a *Application) disconnectFromDB() {
	dbDisconnectCtx, cancel := context.WithTimeout(context.Background(), a.appCtx.Cfg.DBDisconnectTimeout)
	defer cancel()
	err := a.dbClient.Disconnect(dbDisconnectCtx)
	if err != nil {
		log.Info().Printf("failed to disconnect from db: %s\n", err)
	}
	log.Info().Println("Disconnected from DB")
}

func (a *Application) shutdownServer() {
	ctx, cancel := context.WithTimeout(context.Background(), a.appCtx.Cfg.ServerShutdownTimeout)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		log.Error().Fatalf("failed to shutdown http server: %s\n", err)
	}
	log.Info().Println("Server shutdown")
}
