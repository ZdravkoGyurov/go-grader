package docker

import (
	"fmt"
	"log"
)

// ExecuteTests ...
func ExecuteTests(imageName, containerName string) (string, error) {
	err := BuildAssignmentImage("ZdravkoGyurov", "docker-tests", "assignment1", imageName)
	if err != nil {
		return "", fmt.Errorf("failed docker build: %w", err)
	}
	defer handleRemoveImage(imageName)

	output, err := RunImage(imageName, containerName)
	if err != nil && false { // TODO: check additionally if error was from test fail
		return "", fmt.Errorf("failed docker run: %w", err)
	}
	defer handleRemoveContainer(containerName)

	return output, nil
}

func handleRemoveImage(imageName string) {
	if err := RemoveImage(imageName); err != nil {
		log.Printf("failed docker image rm: %s", err)
	}
}

func handleRemoveContainer(containerName string) {
	if err := RemoveContainer(containerName); err != nil {
		log.Printf("failed docker rm: %s", err)
	}
}
