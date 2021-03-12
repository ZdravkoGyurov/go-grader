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
	// create db connection
	dbClient, err := db.Connect(appCtx)
	if err != nil {
		log.Error().Fatalf("failed to connect to mongodb: %s", err)
	}
	log.Info().Println("connected to mongodb...")
	log.Debug().Println("connected to mongodb...")
	log.Error().Println("connected to mongodb...")
	log.Warning().Println("connected to mongodb...")

	// create executor
	exec := executor.New(appCtx.Cfg)
	exec.Start()
	log.Info().Println("Started job executor...")

	// create db handlers
	assignmentsDBHandler := db.NewAssignmentsHandler(dbClient)

	// create http handlers
	assignmentsHTTPHandler := api.NewAssignmentsHandler(assignmentsDBHandler)
	testRunHTTPHandler := api.NewTestRunHandler(exec)

	// create http router
	httpRouter := router.New(assignmentsHTTPHandler, testRunHTTPHandler)

	app := app.New(appCtx, exec, dbClient, httpRouter)

	app.Start()
}
