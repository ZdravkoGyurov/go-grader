package docker

import (
	"fmt"
	"os/exec"
	"syscall"
)

// BuildAssignmentImage ...
func BuildAssignmentImage(gitUser, gitRepo, assignment, imageName string) error {
	cmd := exec.Command("docker", "build",
		"-t", imageName,
		"--build-arg", fmt.Sprintf("gitUser=%s", gitUser),
		"--build-arg", fmt.Sprintf("gitRepo=%s", gitRepo),
		"--build-arg", fmt.Sprintf("assignment=%s", assignment),
		".")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	_, err := cmd.CombinedOutput()
	return err
}

// RunImage runs docker container with given imageName and containerName
func RunImage(imageName, containerName string) (string, error) {
	cmd := exec.Command("docker", "run", "--name", containerName, imageName)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	bytes, err := cmd.CombinedOutput()
	return string(bytes), err
}

// RemoveContainer removes container with given name
func RemoveContainer(containerName string) error {
	cmd := exec.Command("docker", "rm", containerName)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	_, err := cmd.CombinedOutput()
	return err
}

// RemoveImage removes image with given name
func RemoveImage(imageName string) error {
	cmd := exec.Command("docker", "image", "rm", imageName)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	_, err := cmd.CombinedOutput()
	return err
}
