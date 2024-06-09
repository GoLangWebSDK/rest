package rest

import (
	"net/http"
)

type Route struct {
	SubRouter  *http.ServeMux
	Handler    RestHandler
	Host       string
	Schema     string
	Path       string
	PathPrefix string
	FinalPath  string
}

func NewRoute(router *Router) *Route {
	return &Route{
		SubRouter:  router.Mux,
		Handler:    router.CurrentHandler,
		Path:       router.CurrentPath,
		PathPrefix: router.CurrentPathPrefix,
		Host:       router.CurrentHost,
		Schema:     router.CurrentScheme,
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

func (route *Route) SetSchema(schema string) *Route {
	route.Schema = schema
	return route
}

func (route *Route) SetHost(host string) *Route {
	route.Host = host
	return route
}

func (route *Route) Map() {
	route.FinalPath = route.Path

	if route.PathPrefix != "" {
		route.FinalPath = route.PathPrefix + route.Path
	}

	route.Handler.Run()

	// POST
	PostHandler := func(w http.ResponseWriter, r *http.Request) {
		route.Handler.Create(NewSession(r, w))
	}

	PostPath := "POST " + route.FinalPath
	route.SubRouter.HandleFunc(PostPath, PostHandler)

	// GET
	GetHandler := func(w http.ResponseWriter, r *http.Request) {
		route.Handler.Read(NewSession(r, w))
	}

	GetPath := "GET " + route.FinalPath + "/{id}"
	route.SubRouter.HandleFunc(GetPath, GetHandler)

	// GET ALL
	GetAllHandler := func(w http.ResponseWriter, r *http.Request) {
		route.Handler.ReadAll(NewSession(r, w))
	}

	GetAllPath := "GET " + route.FinalPath
	route.SubRouter.HandleFunc(GetAllPath, GetAllHandler)

	// PUT
	PutHandler := func(w http.ResponseWriter, r *http.Request) {
		route.Handler.Update(NewSession(r, w))
	}

	PutPath := "PUT " + route.FinalPath + "/{id}"
	route.SubRouter.HandleFunc(PutPath, PutHandler)

	// DELETE
	DeleteHandler := func(w http.ResponseWriter, r *http.Request) {
		route.Handler.Destroy(NewSession(r, w))
	}

	DeletePath := "DELETE " + route.FinalPath + "/{id}"
	route.SubRouter.HandleFunc(DeletePath, DeleteHandler)
}

// func (rest *Rest) mapRoute() {
// 	pathPrefix := rest.currentRoute
//
// 	if rest.currentRoutePrefix != "" {
// 		pathPrefix = rest.currentRoutePrefix + rest.currentRoute
// 	}
//
// 	route := rest.Mux.PathPrefix(pathPrefix)
//
// 	if rest.currentScheme != "" {
// 		route = route.Schemes(rest.currentScheme)
// 	}
//
// 	if rest.currentHost != "" {
// 		route = route.Host(rest.currentHost)
// 	}
//
// 	rest.SubRouter = route.Subrouter()
// }
//
// func (rest *Rest) mapControllerHandlers() {
// 	ctrl := NewMappedController(rest)
// 	ctrl.Map()
// }
