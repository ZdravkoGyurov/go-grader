package main

import (
	"context"
	"log"

	"grader/api"
	"grader/api/router"
	"grader/db"
	"grader/executor"
	"grader/server"
)

func main() {
	ctx := context.Background()
	client, err := db.Connect(ctx)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %s", err)
	}
	log.Println("connected to mongodb...")

	assignmentsDBHandler := db.NewAssignmentsHandler(client)
	assignmentsHTTPHandler := api.NewAssignmentsHandler(assignmentsDBHandler)

	exec, stop := executor.New(5)
	defer stop()

	testRunHTTPHandler := api.NewTestRunHandler(exec)

	server := server.New("localhost:8080", router.New(assignmentsHTTPHandler, testRunHTTPHandler), client)
	server.Start()
}

// func runTests(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}

// 	gitUser := "ZdravkoGyurov"
// 	gitRepo := "docker-tests"
// 	assignment := "assignment1"
// 	dockerImageName := "week1-image"
// 	dockerContainerName := "week1-container"

// 	err := createTestImage(gitUser, gitRepo, assignment, dockerImageName)
// 	if err != nil {
// 		log.Println(">>> docker build fail: ", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	testResults, err := runTestImage(dockerImageName, dockerContainerName)
// 	if err != nil {
// 		log.Println(">>> docker run fail: ", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	err = removeTestContainer(dockerContainerName)
// 	if err != nil {
// 		log.Println(">>> docker rm fail: ", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	err = removeTestImage(dockerImageName)
// 	if err != nil {
// 		log.Println(">>> docker image rm fail: ", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "%s", testResults)
// 	w.WriteHeader(http.StatusOK)
// }
