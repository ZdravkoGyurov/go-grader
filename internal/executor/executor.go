package executor

import (
	"errors"
	"sync"

	"github.com/ZdravkoGyurov/go-grader/internal/app/config"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/random"
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
	cfg     config.Config
}

// New creates an executor with workers
func New(cfg config.Config) *Executor {
	e := &Executor{
		wg:      sync.WaitGroup{},
		jobs:    make(chan job, cfg.MaxExecutorConcurrentJobs),
		stopped: false,
		cfg:     cfg,
	}

	return e
}

// Start ...
func (e *Executor) Start() {
	e.wg.Add(e.cfg.MaxExecutorWorkers)
	for i := 0; i < e.cfg.MaxExecutorWorkers; i++ {
		go e.startWorker()
	}
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
		log.Info().Printf("enqueued job %s '%s'\n", j.id, j.name)
		return id, nil
	default:
		return "", errors.New("channel is full")
	}
}

func (e *Executor) startWorker() {
	defer e.wg.Done()

	for job := range e.jobs {
		log.Info().Printf("executing job %s\n", job.id)
		job.f()
		log.Info().Printf("finished job %s\n", job.id)
	}
}

// Stop the executor, waits for all jobs to finish
func (e *Executor) Stop() {
	if e.stopped == false {
		e.stopped = true
		close(e.jobs)
		e.wg.Wait()
	}
}
