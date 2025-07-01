package restserver

import "net/http"

type Server interface {
	RegisterPublicRoute(method, path string, handler http.HandlerFunc)
	RegisterSwaggerRoutes()
	Start(address string) error
}
