package main

import (
	"github.com/ZdravkoGyurov/go-grader/api/handlers"
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

	exec := executor.New(appContext.Cfg)
	exec.Start()
	log.Info().Println("Started job executor...")

	httpHandlers := handlers.NewHandlers(appContext, storage, exec)
	httpRouter := router.New(appContext, storage, httpHandlers)

	app := app.New(appContext, exec, storage, httpRouter)

	app.Start()
}
