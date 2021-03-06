package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const (
	// 5MB
	maxUploadSizeInBytes int64 = 5242880
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading file...")

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSizeInBytes)
	r.ParseMultipartForm(maxUploadSizeInBytes)

	file, fileHandler, err := r.FormFile("file-input-name")
	if err != nil {
		fmt.Fprintf(w, "could not upload file, max file size is %dMB\n", maxUploadSizeInBytes/1024/1024)
		return
	}
	defer file.Close()

	// check if file exists already

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("2", err)
	}

	err = os.WriteFile("calc/"+fileHandler.Filename, fileBytes, 0644)
	if err != nil {
		log.Fatal("3", err)
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")

	fmt.Println("Running tests...")
	// decompress archive
	dockerImageName := "calc-tests-img"
	dockerContainerName := "calc-tests-container"

	fmt.Println("Creating test image...")
	err = createTestImage(dockerImageName)
	if err != nil {
		fmt.Println(">>> docker build fail: ", err)
		return
	}

	defer func() {
		fmt.Println("Removing test image...")
		err := removeTestImage(dockerImageName)
		if err != nil {
			fmt.Println(">>> docker image rm fail: ", err)
			return
		}
	}()

	fmt.Println("Running test container...")
	testResults, err := runTestImage(dockerImageName, dockerContainerName)
	if err != nil {
		fmt.Println(">>> docker run fail: ", err)
	}
	fmt.Println("tests run results: ", testResults)

	fmt.Println("Removing test container...")
	err = removeTestContainer(dockerContainerName)
	if err != nil {
		fmt.Println(">>> docker rm fail: ", err)
		return
	}
}

func createTestImage(imageName string) error {
	_, err := exec.Command("docker", "build", "-t", imageName, ".").CombinedOutput()
	return err
}

func runTestImage(imageName, containerName string) (string, error) {
	bytes, err := exec.Command("docker", "run", "--name", containerName, imageName).CombinedOutput()
	return string(bytes), err
}

func removeTestContainer(containerName string) error {
	_, err := exec.Command("docker", "rm", containerName).CombinedOutput()
	return err
}

func removeTestImage(imageName string) error {
	_, err := exec.Command("docker", "image", "rm", imageName).CombinedOutput()
	return err
}

func main() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}
