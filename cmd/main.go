package main

import (
	"log"
	"time"

	"grader/api"
	"grader/api/router"
	"grader/app"
	"grader/db"
	"grader/executor"
)

func main() {
	appCtx := app.NewContext()
	// create db connection
	dbClient, err := db.Connect(appCtx)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %s", err)
	}
	log.Println("connected to mongodb...")

	// create executor
	exec := executor.New(appCtx.Cfg)
	exec.Start()
	exec.EnqueueJob("wait", func() {
		time.Sleep(10 * time.Second)
	})
	log.Println("Started job executor...")

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
