package main

import (
	"github.com/ZdravkoGyurov/go-grader/api/handlers"
	"github.com/ZdravkoGyurov/go-grader/api/middlewares"
	"github.com/ZdravkoGyurov/go-grader/api/router"
	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/executor"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/storage"
)

func main() {
	appContext := app.NewContext()

	storage, err := storage.New(appContext.Context, appContext.Cfg)
	if err != nil {
		log.Error().Fatalf("failed to connect to mongodb: %s", err)
	}
	log.Info().Println("connected to mongodb...")

	exe := executor.New(appContext.Cfg)
	exe.Start()
	log.Info().Println("Started job executor...")

	httpHandlers := handlers.NewHandlers(appContext, storage, exe)
	httpMiddlewares := middlewares.NewMiddlewares(appContext, storage)
	httpRouter := router.New(appContext, httpHandlers, httpMiddlewares)

	app := app.New(appContext, exe, storage, httpRouter)

	app.Start()
}
