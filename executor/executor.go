package executor

import (
	"errors"
	"grader/random"
	"log"
	"sync"
)

const (
	maxConcurrentJobs = 100
)

type job struct {
	id   string
	name string
	f    func()
}

type executor struct {
	wg      sync.WaitGroup
	jobs    chan job
	stopped bool
}

// New creates an executor with workers
func New(workers int) (e *executor, stop func()) {
	e = &executor{
		wg:      sync.WaitGroup{},
		jobs:    make(chan job, maxConcurrentJobs),
		stopped: false,
	}
	e.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go e.startWorker()
	}

	return e, e.stop
}

// QueueJob ...
func (e *executor) EnqueueJob(name string, f func()) (id string, err error) {
	if e.stopped {
		return "", errors.New("executor is stopped")
	}

	id = random.String()
	j := job{
		id:   id,
		name: name,
		f:    f,
	}

	select {
	case e.jobs <- j:
		log.Printf("enqueued job %s '%s'\n", j.id, j.name)
		return id, nil
	default:
		return "", errors.New("channel is full")
	}
}

func (e *executor) startWorker() {
	defer e.wg.Done()

	for job := range e.jobs {
		log.Printf("executing job %s\n", job.id)
		job.f()
		log.Printf("finished job %s\n", job.id)
	}
}

func (e *executor) stop() {
	e.stopped = true
	close(e.jobs)
	e.wg.Wait()
}
