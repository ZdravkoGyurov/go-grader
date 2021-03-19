package dexec

import (
	"fmt"
	"os/exec"
	"syscall"
)

func buildAssignmentImage(testsConfig TestsRunConfig, dockerfile string) (string, error) {
	cmd := exec.Command("docker", "build",
		"--no-cache",
		"-t", testsConfig.ImageName,
		"--build-arg", fmt.Sprintf("assignment=%s", testsConfig.Assignment),
		"--build-arg", fmt.Sprintf("solutionGitUser=%s", testsConfig.SolutionGitUser),
		"--build-arg", fmt.Sprintf("solutionGitRepo=%s", testsConfig.SolutionGitRepo),
		"--build-arg", fmt.Sprintf("testsGitUser=%s", testsConfig.TestsGitUser),
		"--build-arg", fmt.Sprintf("testsGitRepo=%s", testsConfig.TestsGitRepo),
		dockerfile)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	output, err := cmd.CombinedOutput()
	return string(output), err
}

func runImage(imageName, containerName string) (string, error) {
	cmd := exec.Command("docker", "run", "--name", containerName, imageName)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	bytes, err := cmd.CombinedOutput()
	return string(bytes), err
}

func removeContainer(containerName string) error {
	cmd := exec.Command("docker", "rm", containerName)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	_, err := cmd.CombinedOutput()
	return err
}

func removeImage(imageName string) error {
	cmd := exec.Command("docker", "image", "rm", imageName)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	_, err := cmd.CombinedOutput()
	return err
}
