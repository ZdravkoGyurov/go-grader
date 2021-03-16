package docker

import (
	"fmt"

	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/docker/docker/client"
)

type ExecuteTestsConfig struct {
	ImageName       string
	ContainerName   string
	Assignment      string
	SolutionGitUser string
	SolutionGitRepo string
	TestsGitUser    string
	TestsGitRepo    string
}

// ExecuteTests ...
func ExecuteTests(testsCfg ExecuteTestsConfig) (string, error) {
	client.NewEnvClient()
	output, err := BuildAssignmentImage(testsCfg)
	if err != nil {
		return output, fmt.Errorf("failed docker build: %w", err)
	}
	defer handleRemoveImage(testsCfg.ImageName)

	output, err = RunImage(testsCfg.ImageName, testsCfg.ContainerName)
	if err != nil && false { // TODO: check additionally if error was from test fail
		return output, fmt.Errorf("failed docker run: %w", err)
	}
	defer handleRemoveContainer(testsCfg.ContainerName)

	return output, nil
}

func handleRemoveImage(imageName string) {
	if err := RemoveImage(imageName); err != nil {
		log.Info().Printf("failed docker image rm: %s", err)
	}
}

func handleRemoveContainer(containerName string) {
	if err := RemoveContainer(containerName); err != nil {
		log.Info().Printf("failed docker rm: %s", err)
	}
}
