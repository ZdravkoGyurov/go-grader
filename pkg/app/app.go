package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/router"
	"github.com/ZdravkoGyurov/go-grader/pkg/app/config"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
	"github.com/ZdravkoGyurov/go-grader/pkg/executor"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/storage"
)

type Context struct {
	Context context.Context
	Cancel  context.CancelFunc
	Config  config.Config
}

func NewContext() Context {
	ctx, cancel := context.WithCancel(context.Background())
	return Context{
		Context: ctx,
		Cancel:  cancel,
		Config:  config.DefaultConfig(),
	}
}

type Application struct {
	appContext Context
	exe        *executor.Executor
	storage    *storage.Storage
	server     *http.Server
}

func New() *Application {
	appContext := NewContext()

	storage, err := storage.New(appContext.Context, appContext.Config)
	if err != nil {
		log.Error().Fatalf("failed to connect to mongodb: %s", err)
	}
	log.Info().Println("connected to mongodb...")

	exe := executor.New(appContext.Config)
	exe.Start()
	log.Info().Println("started job executor...")

	ctrl := controller.New(appContext.Config, storage, exe)
	httpRouter := router.New(ctrl)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", appContext.Config.Host, appContext.Config.Port),
		Handler:      httpRouter,
		ReadTimeout:  appContext.Config.ServerReadTimeout,
		WriteTimeout: appContext.Config.ServerWriteTimeout,
	}

	return &Application{
		appContext: appContext,
		exe:        exe,
		storage:    storage,
		server:     server,
	}
}

func (a *Application) Start() {
	a.setupSignalNotifier()

	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Fatalf("failed listen and serve: %s\n", err)
		}
	}()
	log.Info().Println("application started...")

	<-a.appContext.Context.Done()

	a.stopExecutor()
	a.disconnectFromDB()
	a.shutdownServer()
	log.Info().Println("Application stopped gracefully")
}

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
	a.exe.Stop()
	log.Info().Println("Executor stopped")
}

func (a *Application) disconnectFromDB() {
	a.storage.Disconnect()
	log.Info().Println("Disconnected from DB")
}

func (a *Application) shutdownServer() {
	ctx, cancel := context.WithTimeout(context.Background(), a.appContext.Config.ServerShutdownTimeout)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		log.Error().Fatalf("failed to shutdown http server: %s\n", err)
	}
	log.Info().Println("Server shutdown")
}
