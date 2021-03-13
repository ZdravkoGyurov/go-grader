package main

import (
	"grader/api"
	"grader/api/router"
	"grader/app"
	"grader/db"
	"grader/executor"
	"grader/log"
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
	httpRouter := router.New(httpHandler)

	app := app.New(appCtx, exec, dbClient, httpRouter)

	app.Start()
}
