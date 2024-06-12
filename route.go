package rest

import (
	"net/http"
)

type Route struct {
	SubRouter   *http.ServeMux
	Handler     RestHandler
	Path        string
	PathPrefix  string
	FinalPath   string
	mappedPaths []string
}

func NewRoute(router *Router) *Route {
	return &Route{
		SubRouter:  router.Mux,
		Handler:    router.CurrentHandler,
		Path:       router.CurrentPath,
		PathPrefix: router.CurrentPathPrefix,
		FinalPath:  router.CurrentPath,
	}
}

func (route *Route) SetPath(path string) *Route {
	route.Path = path
	return route
}

func (route *Route) SetHandler(ctrl RestHandler) *Route {
	route.Handler = ctrl
	return route
}

func (route *Route) Map() {
	if route.PathPrefix != "" {
		route.FinalPath = route.PathPrefix + route.Path
	}

	ctrl := route.Handler.Run()
	mappedPaths := ctrl.Router.CurrentRoute.mappedPaths

	// POST
	PostPath := "POST " + route.FinalPath

	if !route.isRouteMapped(PostPath, mappedPaths...) {
		PostHandler := func(w http.ResponseWriter, r *http.Request) {
			route.Handler.Create(NewSession(r, w))
		}
		route.SubRouter.HandleFunc(PostPath, PostHandler)
	}

	// GET
	GetPath := "GET " + route.FinalPath + "/{id}"

	if !route.isRouteMapped(GetPath, mappedPaths...) {
		GetHandler := func(w http.ResponseWriter, r *http.Request) {
			route.Handler.Read(NewSession(r, w))
		}
		route.SubRouter.HandleFunc(GetPath, GetHandler)
	}

	// GET ALL
	GetAllPath := "GET " + route.FinalPath

	if !route.isRouteMapped(GetAllPath, mappedPaths...) {
		GetAllHandler := func(w http.ResponseWriter, r *http.Request) {
			route.Handler.ReadAll(NewSession(r, w))
		}
		route.SubRouter.HandleFunc(GetAllPath, GetAllHandler)
	}

	// PUT
	PutPath := "PUT " + route.FinalPath + "/{id}"

	if !route.isRouteMapped(PutPath, mappedPaths...) {
		PutHandler := func(w http.ResponseWriter, r *http.Request) {
			route.Handler.Update(NewSession(r, w))
		}
		route.SubRouter.HandleFunc(PutPath, PutHandler)
	}

	// DELETE
	DeletePath := "DELETE " + route.FinalPath + "/{id}"

	if !route.isRouteMapped(DeletePath, mappedPaths...) {
		DeleteHandler := func(w http.ResponseWriter, r *http.Request) {
			route.Handler.Destroy(NewSession(r, w))
		}
		route.SubRouter.HandleFunc(DeletePath, DeleteHandler)
	}
}

func (route *Route) isRouteMapped(path string, mappedRoutes ...string) bool {
	for _, mappedRoute := range mappedRoutes {
		if path == mappedRoute {
			return true
		}
	}
	return false
}
