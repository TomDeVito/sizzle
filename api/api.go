package api

import "net/http"

type APIHandler interface {
	LoadAPIHandler() http.Handler
	GetAPIHandlerParentRoute() string
}
