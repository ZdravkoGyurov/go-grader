package req

import (
	"context"
	"net/http"
)

type reqDataKey struct{}

var requestDataKey reqDataKey

type Data struct {
	Permissions    []string
	GithubUsername string
}

func AddRequestData(r *http.Request, data Data) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), requestDataKey, data))
}

func GetRequestData(r *http.Request) (Data, bool) {
	data, ok := r.Context().Value(requestDataKey).(Data)
	return data, ok
}
