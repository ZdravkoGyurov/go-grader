package middlewares

import (
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/req"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
)

const (
	CreateAssignmentPermission = "CREATE_ASSIGNMENT"
	ReadAssignmentPermission   = "READ_ASSIGNMENT"
	UpdateAssignmentPermission = "UPDATE_ASSIGNMENT"
	DeleteAssignmentPermission = "DELETE_ASSIGNMENT"

	CreateCoursePermission = "CREATE_COURSE"
	ReadCoursePermission   = "READ_COURSE"
	UpdateCoursePermission = "UPDATE_COURSE"
	DeleteCoursePermission = "DELETE_COURSE"

	CreateSubmissionPermission = "CREATE_SUBMISSION"

	ReadUsersPermission  = "READ_USERS"
	UpdateUserPermission = "UPDATE_USER"
	DeleteUserPermission = "DELETE_USER"
)

type Authorization struct {
	RequiredPermissions []string
}

func (m Authorization) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		userPermissions, ok := req.GetPermissions(request)
		if !ok {
			err := errors.New("failed to get user from request context values")
			response.SendError(writer, http.StatusInternalServerError, err)
			return
		}

		userPermMap := permissionsMap(userPermissions...)
		for _, perm := range m.RequiredPermissions {
			if _, ok := userPermMap[perm]; !ok {
				err := errors.Newf("failed to authorize user, missing %s permission", perm)
				response.SendError(writer, http.StatusForbidden, err)
				return
			}
		}

		next.ServeHTTP(writer, request)
	})
}

func permissionsMap(permissions ...string) map[string]struct{} {
	m := make(map[string]struct{}, len(permissions))
	for _, p := range permissions {
		m[p] = struct{}{}
	}
	return m
}
