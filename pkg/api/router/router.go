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
	setupCourseRoutes(r, ctrl)
	setupSubmissionRoutes(r, ctrl)
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
		HandleFunc(paths.Assignment, assignmentHandler.GetAll).Methods(http.MethodGet)

	authRouter(r, ctrl, middlewares.ReadAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, assignmentHandler.Get).Methods(http.MethodGet)

	authRouter(r, ctrl, middlewares.UpdateAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, assignmentHandler.Patch).Methods(http.MethodPatch)

	authRouter(r, ctrl, middlewares.DeleteAssignmentPermission).
		HandleFunc(paths.AssignmentWithID, assignmentHandler.Delete).Methods(http.MethodDelete)
}

func setupCourseRoutes(r *mux.Router, ctrl *controller.Controller) {
	courseHandler := &handlers.Course{Controller: ctrl}
	authRouter(r, ctrl, middlewares.CreateCoursePermission).
		HandleFunc(paths.Course, courseHandler.Post).Methods(http.MethodPost)

	authRouter(r, ctrl, middlewares.ReadCoursePermission).
		HandleFunc(paths.Course, courseHandler.GetAll).Methods(http.MethodGet)

	authRouter(r, ctrl, middlewares.ReadCoursePermission).
		HandleFunc(paths.CourseWithID, courseHandler.Get).Methods(http.MethodGet)

	authRouter(r, ctrl, middlewares.UpdateCoursePermission).
		HandleFunc(paths.CourseWithID, courseHandler.Patch).Methods(http.MethodPatch)

	authRouter(r, ctrl, middlewares.DeleteAssignmentPermission).
		HandleFunc(paths.CourseWithID, courseHandler.Delete).Methods(http.MethodDelete)
}

func setupSubmissionRoutes(r *mux.Router, ctrl *controller.Controller) {
	submissionHandler := &handlers.Submission{Controller: ctrl}
	authRouter(r, ctrl, middlewares.CreateSubmissionPermission).
		HandleFunc(paths.Submission, submissionHandler.Post).Methods(http.MethodPost)
}

func authRouter(r *mux.Router, ctrl *controller.Controller, requiredPermissions ...string) *mux.Router {
	authSubrouter := r.NewRoute().Subrouter()
	authSubrouter.Use(
		middlewares.Authentication{Controller: ctrl}.Authenticate)
	authSubrouter.Use(
		middlewares.Authorization{RequiredPermissions: requiredPermissions}.Authorize)
	return authSubrouter
}
