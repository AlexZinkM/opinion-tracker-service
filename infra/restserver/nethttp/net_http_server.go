package nethttp

import (
	"net/http"

	"opinion-tracker-service/adapter"
	"opinion-tracker-service/infra/restserver"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

type NetHttpServer struct {
	routes []route
}

func NewNetHttpServer() restserver.Server {
	return &NetHttpServer{
		routes: make([]route, 0),
	}
}

func (s *NetHttpServer) RegisterPublicRoute(method, path string, handler http.HandlerFunc) {
	wrappedHandler := adapter.LoggingMiddleware(handler)
	s.routes = append(s.routes, route{
		method:  method,
		path:    path,
		handler: wrappedHandler,
	})
}

func (s *NetHttpServer) RegisterSwaggerRoutes() {
	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)

	http.Handle("/swagger/", swaggerHandler)
}

func (s *NetHttpServer) Start(address string) error {
	http.HandleFunc("/", s.handleAll)
	return http.ListenAndServe(address, nil)
}

func (s *NetHttpServer) handleAll(w http.ResponseWriter, r *http.Request) {
	for _, route := range s.routes {
		if route.method == r.Method && route.path == r.URL.Path {
			route.handler(w, r)
			return
		}
	}

}
