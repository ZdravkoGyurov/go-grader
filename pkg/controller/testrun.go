package controller

import (
	"github.com/ZdravkoGyurov/go-grader/pkg/dexec"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/random"
)

func (c *Controller) CreateTestrun() (string, error) {
	jobName := "run tests in docker"
	jobFunc := func() {
		testsCfg := dexec.TestsRunConfig{
			ImageName:       random.LowercaseString(10),
			ContainerName:   random.LowercaseString(10),
			Assignment:      "assignment1",             // get from body
			SolutionGitUser: "ZdravkoGyurov",           // get from db/user
			SolutionGitRepo: "grader-docker-solutions", // get from db/course
			TestsGitUser:    c.Config.TestsGitUser,
			TestsGitRepo:    c.Config.TestsGitRepo,
		}
		output, err := dexec.RunTests(testsCfg)
		if err != nil {
			log.Error().Println(err) // log status in db
			log.Error().Println(output)
			return
		}
		log.Info().Println(">>>", output) // log status in db
	}
	jobID, err := c.Executor.QueueJob(jobName, jobFunc)
	if err != nil {
		return "", err
	}

	return jobID, nil
}
