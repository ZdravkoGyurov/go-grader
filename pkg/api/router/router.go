package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/handlers"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/middlewares"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/router/paths"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
)

// New creates a mux router with configured routes
func New(ctrl *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.PanicRecovery)
	setupAccountRoutes(r, ctrl)
	setupAssignmentRoutes(r, ctrl)
	setupTestRunRoutes(r, ctrl)
	return r
}

func setupAccountRoutes(r *mux.Router, ctrl *controller.Controller) {
	registrationHandler := &handlers.Registration{Controller: ctrl}
	r.HandleFunc(paths.Register, registrationHandler.Post).Methods(http.MethodPost)

	loginHandler := &handlers.Login{Controller: ctrl}
	r.HandleFunc(paths.Login, loginHandler.Post).Methods(http.MethodPost)

	logoutHandler := &handlers.Logout{Controller: ctrl}
	r.HandleFunc(paths.Logout, logoutHandler.Post).Methods(http.MethodPost)
}

func setupAssignmentRoutes(r *mux.Router, ctrl *controller.Controller) {
	assignmentHandler := &handlers.Assignment{Controller: ctrl}
	authRouter(r, ctrl, middlewares.CreateAssignmentPermission).
		HandleFunc(paths.Assignment, assignmentHandler.Post).Methods(http.MethodPost)

	authRouter(r, ctrl, middlewares.ReadAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, assignmentHandler.Get).Methods(http.MethodGet)

	authRouter(r, ctrl, middlewares.UpdatessignmentPermission).
		HandleFunc(paths.AssignmentWithID, assignmentHandler.Patch).Methods(http.MethodPatch)

	authRouter(r, ctrl, middlewares.DeleteAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, assignmentHandler.Delete).Methods(http.MethodDelete)
}

func setupTestRunRoutes(r *mux.Router, ctrl *controller.Controller) {
	testrunHandler := &handlers.Testrun{Controller: ctrl}
	authRouter(r, ctrl, middlewares.CreateTestRunPermission).
		HandleFunc(paths.TestRun, testrunHandler.Post).Methods(http.MethodPost)
}

func authRouter(r *mux.Router, ctrl *controller.Controller, requiredPermissions ...string) *mux.Router {
	authSubrouter := r.NewRoute().Subrouter()
	authSubrouter.Use(
		middlewares.Authentication{Controller: ctrl}.Authenticate)
	authSubrouter.Use(
		middlewares.Authorization{RequiredPermissions: requiredPermissions}.Authorize)
	return authSubrouter
}
