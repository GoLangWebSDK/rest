package rest

//ToDo
// - add strict slash support

import (
	"net/http"
)

type Router struct {
	Mux               *http.ServeMux
	HTTPHandler       http.Handler
	CurrentRoute      *Route
	CurrentPath       string
	CurrentPathPrefix string
	CurrentHandler    RestHandler
}

func NewRouter() *Router {
	return &Router{
		Mux: http.NewServeMux(),
	}
}

func (router *Router) Load(routes Routes) *Router {
	routes.LoadRoutes(router)
	routes.LoadMiddleware(router)
	return router
}

func (router *Router) Use(middlewares ...Middleware) *Router {
	if router.HTTPHandler == nil {
		router.HTTPHandler = router.Mux
	}

	for _, mw := range middlewares {
		router.HTTPHandler = mw(router.HTTPHandler)
	}

	return router
}

func (router *Router) RoutePrefix(prefix string) *Router {
	router.CurrentPathPrefix = prefix
	return router
}

func (router *Router) API(version ...string) *Router {
	if len(version) != 0 {
		router.CurrentPathPrefix = "/api/" + version[0]
		return router
	}
	router.CurrentPathPrefix = "/api"
	return router
}

func (router *Router) Route(path string) *Router {
	router.CurrentPath = path
	router.CurrentRoute = NewRoute(router)
	return router
}

func (router *Router) Controller(ctrl RestHandler) {
	router.CurrentRoute.SetHandler(ctrl)
	router.CurrentHandler = ctrl

	router.CurrentRoute.Map()
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if router.HTTPHandler == nil {
		router.HTTPHandler = router.Mux
	}

	router.HTTPHandler.ServeHTTP(w, req)
}
