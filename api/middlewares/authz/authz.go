package authz

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/api/req"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

const (
	CreateAssignmentPermission = "CREATE_ASSIGNMENT"
	ReadAssignmentPermission   = "READ_ASSIGNMENT"
	UpdatessignmentPermission  = "UPDATE_ASSIGNMENT"
	DeleteAssignmentPermission = "DELETE_ASSIGNMENT"

	CreateTestRunPermission = "CREATE_TESTRUN"
)

type middleware struct {
	requiredPermissions []string
}

func Middleware(requiredPermissions ...string) func(http.Handler) http.Handler {
	mw := &middleware{
		requiredPermissions: requiredPermissions,
	}
	return mw.authorize
}

func (m middleware) authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		userPermissions, ok := req.GetPermissions(request)
		if !ok {
			log.Error().Println(errors.New("failed to get user from request context values"))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		userPermMap := permissionsMap(userPermissions...)
		for _, perm := range m.requiredPermissions {
			if _, ok := userPermMap[perm]; !ok {
				log.Error().Println(fmt.Errorf("failed to authorize user, missing %s permission", perm))
				writer.WriteHeader(http.StatusForbidden)
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
