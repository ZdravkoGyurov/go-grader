package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/handlers"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/middlewares"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/router/paths"
	"github.com/ZdravkoGyurov/go-grader/pkg/app"
)

// New creates a mux router with configured routes
func New(appContext app.Context, handlers *handlers.Handlers, mws *middlewares.Middlewares) *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.PanicRecovery)
	setupAccountRoutes(r, appContext, handlers)
	setupAssignmentRoutes(r, appContext, handlers, mws)
	setupTestRunRoutes(r, appContext, handlers, mws)
	return r
}

func setupAccountRoutes(r *mux.Router, appContext app.Context, httpHandlers *handlers.Handlers) {
	r.HandleFunc(paths.Register, httpHandlers.Registration.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.Login, httpHandlers.Login.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.Logout, httpHandlers.Logout.Post).Methods(http.MethodPost)
}

func setupAssignmentRoutes(r *mux.Router, appContext app.Context, httpHandlers *handlers.Handlers, mws *middlewares.Middlewares) {
	authRouter(r, appContext, mws, middlewares.CreateAssignmentPermission).
		HandleFunc(paths.Assignment, httpHandlers.Assignment.Post).Methods(http.MethodPost)

	authRouter(r, appContext, mws, middlewares.ReadAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, httpHandlers.Assignment.Get).Methods(http.MethodGet)

	authRouter(r, appContext, mws, middlewares.UpdatessignmentPermission).
		HandleFunc(paths.AssignmentWithID, httpHandlers.Assignment.Patch).Methods(http.MethodPatch)

	authRouter(r, appContext, mws, middlewares.DeleteAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, httpHandlers.Assignment.Delete).Methods(http.MethodDelete)
}

func setupTestRunRoutes(r *mux.Router, appContext app.Context, httpHandlers *handlers.Handlers, mws *middlewares.Middlewares) {
	authRouter(r, appContext, mws, middlewares.CreateTestRunPermission).
		HandleFunc(paths.TestRun, httpHandlers.TestRun.Post).Methods(http.MethodPost)
}

func authRouter(r *mux.Router, appContext app.Context, mws *middlewares.Middlewares, requiredPermissions ...string) *mux.Router {
	authSubrouter := r.NewRoute().Subrouter()
	mws.ApplyAll(authSubrouter)
	authzMiddleware := middlewares.AuthzMiddleware{
		RequiredPermissions: requiredPermissions,
	}
	authSubrouter.Use(authzMiddleware.Authorize)
	return authSubrouter
}
