package rest

import (
	"net/http"
)

type Rest struct {
	Mux               *http.ServeMux
	HTTPHandler       http.Handler
	CurrentRoute      *Route
	CurrentScheme     string
	CurrentHost       string
	CurrentPathPrefix string
	CurrentPath       string
	currentHandler    RestHandler
}

func NewRouter() *Rest {
	return &Rest{
		Mux: http.NewServeMux(),
	}
}

func (rest *Rest) Load(routes Routes) {
	if routes != nil {
		routes.LoadRoutes(rest)
		routes.LoadMiddlewares(rest)
	}
}

func (rest *Rest) Use(middlewares ...Middleware) *Rest {
	if rest.HTTPHandler == nil {
		rest.HTTPHandler = rest.Mux
	}

	for _, mw := range middlewares {
		rest.HTTPHandler = mw(rest.HTTPHandler)
	}

	return rest
}

func (rest *Rest) Scheme(scheme string) *Rest {
	rest.CurrentScheme = scheme
	return rest
}

func (rest *Rest) Host(host string) *Rest {
	rest.CurrentHost = host
	return rest
}

func (rest *Rest) RoutePrefix(prefix string) *Rest {
	rest.CurrentPathPrefix = prefix
	return rest
}

func (rest *Rest) API(version ...string) *Rest {
	if len(version) != 0 {
		rest.CurrentPathPrefix = "/api/" + version[0]
		return rest
	}
	rest.CurrentPathPrefix = "/api"
	return rest
}

// func (rest *Rest) StrictSlash(value bool) *Rest {
// 	rest.Mux.StrictSlash(value)
// 	return rest
// }

func (rest *Rest) Route(route string) *Rest {
	rest.CurrentPath = route
	rest.CurrentRoute = NewRoute(rest)
	return rest
}

func (rest *Rest) Controller(ctrl RestHandler) {
	rest.currentHandler = ctrl
}

func (rest *Rest) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if rest.HTTPHandler == nil {
		rest.HTTPHandler = rest.Mux
	}

	rest.HTTPHandler.ServeHTTP(w, req)
}
