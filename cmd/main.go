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

	// create executor
	exec := executor.New(appCtx.Cfg)
	exec.Start()
	log.Info().Println("Started job executor...")

	// create db handlers
	assignmentsDBHandler := db.NewAssignmentsHandler(appCtx, dbClient)
	userDBHandler := db.NewUserHandler(appCtx, dbClient)
	sessionDBHandler := db.NewSessionHandler(appCtx, dbClient)

	// create http handlers
	registrationHTTPHandler := api.NewRegistrationHandler(userDBHandler)
	loginHTTPHandler := api.NewLoginHandler(userDBHandler, sessionDBHandler)
	logoutHTTPHandler := api.NewLogoutHandler(sessionDBHandler)
	assignmentsHTTPHandler := api.NewAssignmentsHandler(assignmentsDBHandler)
	testRunHTTPHandler := api.NewTestRunHandler(exec)

	// create http router
	httpRouter := router.New(router.HTTPHandlers{
		Registration: registrationHTTPHandler,
		Login:        loginHTTPHandler,
		Logout:       logoutHTTPHandler,
		Assignments:  assignmentsHTTPHandler,
		TestRun:      testRunHTTPHandler,
	})

	app := app.New(appCtx, exec, dbClient, httpRouter)

	app.Start()
}
