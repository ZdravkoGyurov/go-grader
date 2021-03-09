package main

import (
	"context"
	"log"

	"grader/api"
	"grader/api/router"
	"grader/app"
	"grader/db"
	"grader/executor"
)

func main() {
	ctx := context.Background()

	// create db connection
	dbClient, err := db.Connect(ctx)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %s", err)
	}
	log.Println("connected to mongodb...")

	// create executor
	exec, stopExecutor := executor.New(5)
	log.Println("started job executor...")

	// create db handlers
	assignmentsDBHandler := db.NewAssignmentsHandler(dbClient)

	// create http handlers
	assignmentsHTTPHandler := api.NewAssignmentsHandler(assignmentsDBHandler)
	testRunHTTPHandler := api.NewTestRunHandler(exec)

	// create http router
	httpRouter := router.New(assignmentsHTTPHandler, testRunHTTPHandler)

	app := app.New(stopExecutor, "localhost:8080", httpRouter, dbClient)

	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	app.Stop()
	// }()

	app.Start()
}
