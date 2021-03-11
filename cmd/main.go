package main

import (
	"log"
	"time"

	"grader/api"
	"grader/api/router"
	"grader/app"
	"grader/app/config"
	"grader/db"
	"grader/executor"
)

func main() {
	cfg := config.Config{
		Host:                      "localhost",
		Port:                      8080,
		MaxExecutorWorkers:        5,
		MaxExecutorConcurrentJobs: 100,
		GithubTestsRepo:           "",
		DBConnectTimeout:          30 * time.Second,
	}
	// create db connection
	dbClient, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %s", err)
	}
	log.Println("connected to mongodb...")

	// create executor
	exec, stopExecutor := executor.New(cfg)
	log.Println("started job executor...")

	// create db handlers
	assignmentsDBHandler := db.NewAssignmentsHandler(dbClient)

	// create http handlers
	assignmentsHTTPHandler := api.NewAssignmentsHandler(assignmentsDBHandler)
	testRunHTTPHandler := api.NewTestRunHandler(exec)

	// create http router
	httpRouter := router.New(assignmentsHTTPHandler, testRunHTTPHandler)

	app := app.New(cfg, stopExecutor, httpRouter, dbClient)

	app.Start()
}
