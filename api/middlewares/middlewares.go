package middlewares

import (
	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/storage"
	"github.com/gorilla/mux"
)

type Middlewares struct {
	Authentication *authnMiddleware
}

func NewMiddlewares(appContext app.Context, storage *storage.Storage) *Middlewares {
	return &Middlewares{
		Authentication: &authnMiddleware{appContext: appContext, authnStorage: storage},
	}
}

func (m *Middlewares) ApplyAll(router *mux.Router) {
	router.Use(m.Authentication.authenticate)
}
