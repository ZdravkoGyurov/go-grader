package dexec

import (
	"fmt"

	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

type TestsRunConfig struct {
	ImageName       string
	ContainerName   string
	Assignment      string
	SolutionGitUser string
	SolutionGitRepo string
	TestsGitUser    string
	TestsGitRepo    string
}

func RunTests(testsCfg TestsRunConfig) (string, error) {
	output, err := buildAssignmentImage(testsCfg, "pkg/dexec/.")
	if err != nil {
		return output, fmt.Errorf("failed docker build: %w", err)
	}
	defer handleRemoveImage(testsCfg.ImageName)

	output, err = runImage(testsCfg.ImageName, testsCfg.ContainerName)
	if err != nil && false { // TODO: check additionally if error was from test fail
		return output, fmt.Errorf("failed docker run: %w", err)
	}
	defer handleRemoveContainer(testsCfg.ContainerName)

	return output, nil
}

func handleRemoveImage(imageName string) {
	if err := removeImage(imageName); err != nil {
		log.Error().Printf("failed docker image rm: %s", err)
	}
}

func handleRemoveContainer(containerName string) {
	if err := removeContainer(containerName); err != nil {
		log.Error().Printf("failed docker rm: %s", err)
	}
}
