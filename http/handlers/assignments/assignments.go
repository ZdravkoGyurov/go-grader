package assignments

import (
	"context"
	"net/http"
)

type assignmentsDBHandler interface {
	CreateAssignment(ctx context.Context) error
}

type assignmentsHandler struct {
	dbHandler assignmentsDBHandler
}

// NewHandler creates a new assignments handler
func NewHandler(dbHandler assignmentsDBHandler) *assignmentsHandler {
	return &assignmentsHandler{
		dbHandler: dbHandler,
	}
}

func (h *assignmentsHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
}
