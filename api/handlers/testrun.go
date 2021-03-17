package handlers

import (
	"fmt"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/dexec"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/random"
)

type testRunDBHandler interface {
}

type exe interface {
	QueueJob(name string, f func()) (id string, err error)
}

type testrunHandler struct {
	appContext app.Context
	exe
}

func (h *testrunHandler) Post(writer http.ResponseWriter, request *http.Request) {
	// ctx := request.Context()

	jobName := "run tests in docker"
	jobFunc := func() {
		testsCfg := dexec.TestsRunConfig{
			ImageName:       random.LowercaseString(10),
			ContainerName:   random.LowercaseString(10),
			Assignment:      "assignment1",             // get from body
			SolutionGitUser: "ZdravkoGyurov",           // get from db/user
			SolutionGitRepo: "grader-docker-solutions", // get from db/course
			TestsGitUser:    h.appContext.Cfg.TestsGitUser,
			TestsGitRepo:    h.appContext.Cfg.TestsGitRepo,
		}
		output, err := dexec.RunTests(testsCfg)
		if err != nil {
			log.Error().Println(err) // log status in db
			log.Error().Println(output)
			return
		}
		log.Info().Println(">>>", output) // log status in db
	}
	jobID, err := h.exe.QueueJob(jobName, jobFunc)
	if err != nil {
		log.Error().Printf("failed to run job '%s': %s", jobName, err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf(`{"jobId": %s}`, jobID)))
}
