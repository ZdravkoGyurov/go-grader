package docker

import (
	"fmt"
	"os/exec"
	"syscall"
)

// BuildAssignmentImage ...
func BuildAssignmentImage(testsCfg ExecuteTestsConfig) (string, error) {
	cmd := exec.Command("docker", "build",
		"--no-cache",
		"-t", testsCfg.ImageName,
		"--build-arg", fmt.Sprintf("assignment=%s", testsCfg.Assignment),
		"--build-arg", fmt.Sprintf("solutionGitUser=%s", testsCfg.SolutionGitUser),
		"--build-arg", fmt.Sprintf("solutionGitRepo=%s", testsCfg.SolutionGitRepo),
		"--build-arg", fmt.Sprintf("testsGitUser=%s", testsCfg.TestsGitUser),
		"--build-arg", fmt.Sprintf("testsGitRepo=%s", testsCfg.TestsGitRepo),
		"./docker")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	output, err := cmd.CombinedOutput()
	return string(output), err
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
