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
func New(ctrl *controller.Controller, mws *middlewares.Middlewares) *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.PanicRecovery)
	setupAccountRoutes(r, ctrl)
	setupAssignmentRoutes(r, ctrl, mws)
	setupTestRunRoutes(r, ctrl, mws)
	return r
}

func setupAccountRoutes(r *mux.Router, ctrl *controller.Controller) {
	registrationHandler := &handlers.RegistrationHandler{Controller: ctrl}
	r.HandleFunc(paths.Register, registrationHandler.Post).Methods(http.MethodPost)

	loginHandler := &handlers.LoginHandler{Controller: ctrl}
	r.HandleFunc(paths.Login, loginHandler.Post).Methods(http.MethodPost)

	logoutHandler := &handlers.LogoutHandler{Controller: ctrl}
	r.HandleFunc(paths.Logout, logoutHandler.Post).Methods(http.MethodPost)
}

func setupAssignmentRoutes(r *mux.Router, ctrl *controller.Controller, mws *middlewares.Middlewares) {
	assignmentHandler := &handlers.AssignmentHandler{Controller: ctrl}
	authRouter(r, mws, middlewares.CreateAssignmentPermission).
		HandleFunc(paths.Assignment, assignmentHandler.Post).Methods(http.MethodPost)

	authRouter(r, mws, middlewares.ReadAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, assignmentHandler.Get).Methods(http.MethodGet)

	authRouter(r, mws, middlewares.UpdatessignmentPermission).
		HandleFunc(paths.AssignmentWithID, assignmentHandler.Patch).Methods(http.MethodPatch)

	authRouter(r, mws, middlewares.DeleteAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, assignmentHandler.Delete).Methods(http.MethodDelete)
}

func setupTestRunRoutes(r *mux.Router, ctrl *controller.Controller, mws *middlewares.Middlewares) {
	testrunHandler := &handlers.TestrunHandler{Controller: ctrl}
	authRouter(r, mws, middlewares.CreateTestRunPermission).
		HandleFunc(paths.TestRun, testrunHandler.Post).Methods(http.MethodPost)
}

func authRouter(r *mux.Router, mws *middlewares.Middlewares, requiredPermissions ...string) *mux.Router {
	authSubrouter := r.NewRoute().Subrouter()
	mws.ApplyAll(authSubrouter)
	authzMiddleware := middlewares.AuthzMiddleware{
		RequiredPermissions: requiredPermissions,
	}
	authSubrouter.Use(authzMiddleware.Authorize)
	return authSubrouter
}
