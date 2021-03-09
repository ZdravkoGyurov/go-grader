package docker

// ExecuteTests ...
func ExecuteTests(imageName, containerName string) (string, error) {
	err := BuildAssignmentImage("ZdravkoGyurov", "docker-tests", "assignment2", imageName)
	if err != nil {
		return "", err
	}

	output, err := RunImage(imageName, containerName)
	if err != nil && false { // TODO: check additionally if error was from test fail
		return "", err
	}

	// fix cleanup
	err = RemoveContainer(containerName)
	if err != nil {
		return "", err
	}

	err = RemoveImage(imageName)
	if err != nil {
		return "", err
	}

	return output, nil
}
