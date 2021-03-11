package executor

import (
	"errors"
	"grader/app/config"
	"grader/random"
	"log"
	"sync"
)

// StopFunc ...
type StopFunc = func()

type job struct {
	id   string
	name string
	f    func()
}

// Executor ...
type Executor struct {
	wg      sync.WaitGroup
	jobs    chan job
	stopped bool
}

// New creates an executor with workers
func New(config config.Config) (*Executor, StopFunc) {
	e := &Executor{
		wg:      sync.WaitGroup{},
		jobs:    make(chan job, config.MaxExecutorConcurrentJobs),
		stopped: false,
	}
	e.wg.Add(config.MaxExecutorWorkers)
	for i := 0; i < config.MaxExecutorWorkers; i++ {
		go e.startWorker()
	}

	return e, e.stop
}

// EnqueueJob ...
func (e *Executor) EnqueueJob(name string, f func()) (id string, err error) {
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

func (e *Executor) startWorker() {
	defer e.wg.Done()

	for job := range e.jobs {
		log.Printf("executing job %s\n", job.id)
		job.f()
		log.Printf("finished job %s\n", job.id)
	}
}

func (e *Executor) stop() {
	e.stopped = true
	close(e.jobs)
	e.wg.Wait()
}
