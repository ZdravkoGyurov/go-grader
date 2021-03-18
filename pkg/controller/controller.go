package controller

import (
	"github.com/ZdravkoGyurov/go-grader/pkg/app/config"
	"github.com/ZdravkoGyurov/go-grader/pkg/executor"
	"github.com/ZdravkoGyurov/go-grader/pkg/storage"
)

type Controller struct {
	Config   config.Config
	storage  *storage.Storage
	executor *executor.Executor
}

func New(cfg config.Config, storage *storage.Storage, executor *executor.Executor) *Controller {
	return &Controller{
		Config:   cfg,
		storage:  storage,
		executor: executor,
	}
}
