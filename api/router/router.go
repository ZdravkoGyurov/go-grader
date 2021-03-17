package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ZdravkoGyurov/go-grader/api/handlers"
	"github.com/ZdravkoGyurov/go-grader/api/middlewares/authn"
	"github.com/ZdravkoGyurov/go-grader/api/middlewares/authz"
	"github.com/ZdravkoGyurov/go-grader/api/router/paths"
	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/storage"
)

// New creates a mux router with configured routes
func New(appContext app.Context, storage *storage.Storage, httpHandlers *handlers.Handlers) *mux.Router {
	r := mux.NewRouter()

	setupAccountRoutes(r, appContext, storage, httpHandlers)
	setupAssignmentRoutes(r, appContext, storage, httpHandlers)
	setupTestRunRoutes(r, appContext, storage, httpHandlers)

	return r
}

func setupAccountRoutes(r *mux.Router, appContext app.Context, storage *storage.Storage, httpHandlers *handlers.Handlers) {
	r.HandleFunc(paths.Register, httpHandlers.Registration.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.Login, httpHandlers.Login.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.Logout, httpHandlers.Logout.Post).Methods(http.MethodPost)
}

func setupAssignmentRoutes(r *mux.Router, appContext app.Context, storage *storage.Storage, httpHandlers *handlers.Handlers) {
	authRouter(r, appContext, storage, authz.CreateAssignmentPermission).
		HandleFunc(paths.Assignment, httpHandlers.Assignment.Post).Methods(http.MethodPost)

	authRouter(r, appContext, storage, authz.ReadAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, httpHandlers.Assignment.Get).Methods(http.MethodGet)

	authRouter(r, appContext, storage, authz.UpdatessignmentPermission).
		HandleFunc(paths.AssignmentWithID, httpHandlers.Assignment.Patch).Methods(http.MethodPatch)

	authRouter(r, appContext, storage, authz.DeleteAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, httpHandlers.Assignment.Delete).Methods(http.MethodDelete)
}

func setupTestRunRoutes(r *mux.Router, appContext app.Context, storage *storage.Storage, httpHandlers *handlers.Handlers) {
	authRouter(r, appContext, storage, authz.CreateTestRunPermission).
		HandleFunc(paths.TestRun, httpHandlers.TestRun.Post).Methods(http.MethodPost)
}

func authRouter(r *mux.Router, appContext app.Context, storage *storage.Storage, requiredPermissions ...string) *mux.Router {
	authSubrouter := r.NewRoute().Subrouter()
	authSubrouter.Use(authn.Middleware(appContext, storage))
	authSubrouter.Use(authz.Middleware(requiredPermissions...))
	return authSubrouter
}
