package main

import (
	"fmt"
	"grader/executor"
	"log"
	"net/http"
	"time"
)

// http://www.inanzzz.com/index.php/post/3hut/a-simple-worker-and-work-queue-example-with-golang

const (
	maxWorkers = 5
)

func main() {
	exec, stop := executor.New(maxWorkers)
	defer stop()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		f := func() {
			time.Sleep(2 * time.Second)
		}
		jobID, err := exec.EnqueueJob("docker run", f)
		if err != nil {
			log.Printf("failed to run job: %s\n", err)
			fmt.Fprintf(w, "failed to run job: %s", err)
			return
		}
		fmt.Fprintf(w, "%s", jobID)
	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		go stop()
		fmt.Fprintf(w, "stopped executor, but finishing queued jobs")
	})
	http.ListenAndServe("localhost:8080", nil)
}
