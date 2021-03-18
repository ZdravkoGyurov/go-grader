package controller

import (
	"github.com/ZdravkoGyurov/go-grader/pkg/app/config"
	"github.com/ZdravkoGyurov/go-grader/pkg/executor"
	"github.com/ZdravkoGyurov/go-grader/pkg/storage"
)

type Controller struct {
	Config   config.Config
	Storage  *storage.Storage
	Executor *executor.Executor
}
