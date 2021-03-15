package main

import (
	"github.com/ZdravkoGyurov/go-grader/api"
	"github.com/ZdravkoGyurov/go-grader/api/router"
	"github.com/ZdravkoGyurov/go-grader/app"
	"github.com/ZdravkoGyurov/go-grader/db"
	"github.com/ZdravkoGyurov/go-grader/executor"
	"github.com/ZdravkoGyurov/go-grader/log"
)

func main() {
	appCtx := app.NewContext()

	dbClient, err := db.Connect(appCtx)
	if err != nil {
		log.Error().Fatalf("failed to connect to mongodb: %s", err)
	}
	log.Info().Println("connected to mongodb...")

	exec := executor.New(appCtx.Cfg)
	exec.Start()
	log.Info().Println("Started job executor...")

	dbHandler := db.NewHandlers(appCtx, dbClient)
	httpHandler := api.NewHandlers(appCtx, *dbHandler, exec)
	httpRouter := router.New(appCtx, dbHandler, httpHandler)

	app := app.New(appCtx, exec, dbClient, httpRouter)

	app.Start()
}
