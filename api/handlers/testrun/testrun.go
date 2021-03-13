package testrun

import (
	"fmt"
	"net/http"

	"grader/docker"
	"grader/log"
	"grader/random"
)

type testRunDBHandler interface {
}

type executor interface {
	EnqueueJob(name string, f func()) (id string, err error)
}

// HTTPHandler ...
type HTTPHandler struct {
	executor
}

// NewHTTPHandler creates a new test run http handler
func NewHTTPHandler(executor executor) *HTTPHandler {
	return &HTTPHandler{
		executor: executor,
	}
}

// Post ...
func (h *HTTPHandler) Post(writer http.ResponseWriter, request *http.Request) {
	// ctx := request.Context()

	jobName := "run tests in docker"
	jobFunc := func() {
		imageName := random.String()
		containerName := random.String()
		output, err := docker.ExecuteTests(imageName, containerName)
		if err != nil {
			log.Info().Println(err) // log status in db
			return
		}
		log.Info().Println(">>>", output) // log status in db
	}
	jobID, err := h.executor.EnqueueJob(jobName, jobFunc)
	if err != nil {
		log.Info().Printf("failed to run job '%s': %s", jobName, err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf(`{"jobId": %s}`, jobID)))
}
