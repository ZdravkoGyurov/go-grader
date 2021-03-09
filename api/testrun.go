package api

import (
	"fmt"
	"log"
	"net/http"
)

type testRunDBHandler interface {
}

type executor interface {
	EnqueueJob(name string, f func()) (id string, err error)
}

// TestRunHandler ...
type TestRunHandler struct {
	executor
}

// NewTestRunHandler creates a new test run http handler
func NewTestRunHandler(executor executor) *TestRunHandler {
	return &TestRunHandler{
		executor: executor,
	}
}

// Post ...
func (h *TestRunHandler) Post(writer http.ResponseWriter, request *http.Request) {
	// ctx := request.Context()

	jobName := "run tests"
	jobFunc := func() {
		fmt.Println("simulating execution of tests run")
	}
	jobID, err := h.executor.EnqueueJob(jobName, jobFunc)
	if err != nil {
		log.Printf("failed to run job '%s': %s", jobName, err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf(`{"jobId": %s}`, jobID)))
}
