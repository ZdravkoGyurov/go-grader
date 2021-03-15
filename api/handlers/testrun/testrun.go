package testrun

import (
	"fmt"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/app"
	"github.com/ZdravkoGyurov/go-grader/docker"
	"github.com/ZdravkoGyurov/go-grader/log"
	"github.com/ZdravkoGyurov/go-grader/random"
)

type testRunDBHandler interface {
}

type executor interface {
	EnqueueJob(name string, f func()) (id string, err error)
}

// HTTPHandler ...
type HTTPHandler struct {
	appCtx app.Context
	executor
}

// NewHTTPHandler creates a new test run http handler
func NewHTTPHandler(appCtx app.Context, executor executor) *HTTPHandler {
	return &HTTPHandler{
		appCtx:   appCtx,
		executor: executor,
	}
}

// Post ...
func (h *HTTPHandler) Post(writer http.ResponseWriter, request *http.Request) {
	// ctx := request.Context()

	jobName := "run tests in docker"
	jobFunc := func() {
		testsCfg := docker.ExecuteTestsConfig{
			ImageName:       random.String(),
			ContainerName:   random.String(),
			Assignment:      "assignment1",             // get from body
			SolutionGitUser: "ZdravkoGyurov",           // get from db/user
			SolutionGitRepo: "grader-docker-solutions", // get from db/course
			TestsGitUser:    h.appCtx.Cfg.TestsGitUser,
			TestsGitRepo:    h.appCtx.Cfg.TestsGitRepo,
		}
		output, err := docker.ExecuteTests(testsCfg)
		if err != nil {
			log.Error().Println(err) // log status in db
			log.Error().Println(output)
			return
		}
		log.Info().Println(">>>", output) // log status in db
	}
	jobID, err := h.executor.EnqueueJob(jobName, jobFunc)
	if err != nil {
		log.Error().Printf("failed to run job '%s': %s", jobName, err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf(`{"jobId": %s}`, jobID)))
}
