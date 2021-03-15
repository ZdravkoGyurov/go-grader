package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ZdravkoGyurov/go-grader/api"
	"github.com/ZdravkoGyurov/go-grader/api/middlewares/authn"
	"github.com/ZdravkoGyurov/go-grader/api/middlewares/authz"
	"github.com/ZdravkoGyurov/go-grader/api/router/paths"
	"github.com/ZdravkoGyurov/go-grader/app"
	"github.com/ZdravkoGyurov/go-grader/db"
)

// New creates a mux router with configured routes
func New(appCtx app.Context, dbHandlers *db.Handlers, httpHandlers *api.Handlers) *mux.Router {
	r := mux.NewRouter()

	setupAccountRoutes(r, appCtx, dbHandlers, httpHandlers)
	setupAssignmentRoutes(r, appCtx, dbHandlers, httpHandlers)
	setupTestRunRoutes(r, appCtx, dbHandlers, httpHandlers)

	return r
}

func setupAccountRoutes(r *mux.Router, appCtx app.Context, dbHandlers *db.Handlers, httpHandlers *api.Handlers) {
	r.HandleFunc(paths.Register, httpHandlers.Registration.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.Login, httpHandlers.Login.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.Logout, httpHandlers.Logout.Post).Methods(http.MethodPost)
}

func setupAssignmentRoutes(r *mux.Router, appCtx app.Context, dbHandlers *db.Handlers, httpHandlers *api.Handlers) {
	authRouter(r, appCtx, dbHandlers, authz.CreateAssignmentPermission).
		HandleFunc(paths.Assignment, httpHandlers.Assignment.Post).Methods(http.MethodPost)

	authRouter(r, appCtx, dbHandlers, authz.ReadAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, httpHandlers.Assignment.Get).Methods(http.MethodGet)

	authRouter(r, appCtx, dbHandlers, authz.UpdatessignmentPermission).
		HandleFunc(paths.AssignmentWithID, httpHandlers.Assignment.Patch).Methods(http.MethodPatch)

	authRouter(r, appCtx, dbHandlers, authz.DeleteAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, httpHandlers.Assignment.Delete).Methods(http.MethodDelete)
}

func setupTestRunRoutes(r *mux.Router, appCtx app.Context, dbHandlers *db.Handlers, httpHandlers *api.Handlers) {
	authRouter(r, appCtx, dbHandlers, authz.CreateTestRunPermission).
		HandleFunc(paths.TestRun, httpHandlers.TestRun.Post).Methods(http.MethodPost)
}

func authRouter(r *mux.Router, appCtx app.Context, dbHandlers *db.Handlers, requiredPermissions ...string) *mux.Router {
	authSubrouter := r.NewRoute().Subrouter()
	authSubrouter.Use(authn.Middleware(appCtx, dbHandlers.Session, dbHandlers.User))
	authSubrouter.Use(authz.Middleware(requiredPermissions...))
	return authSubrouter
}
