package rest

import (
	"github.com/gorilla/mux"
)

type Rest struct {
	Mux                *mux.Router
	SubRouter          *mux.Router
	controllers        []MappedController
	currentScheme      string
	currentHost        string
	currentRoutePrefix string
	currentRoute       string
	currentHandler     RestHandler
}

func NewRouter() *Rest {
	m := mux.NewRouter()
	return &Rest{
		Mux:         m,
		SubRouter:   m,
		controllers: []MappedController{},
	}
}

func (rest *Rest) Load(routes Routes) {
	if routes != nil {
		routes.LoadRoutes(rest)
		routes.LoadMiddlewares(rest)
	}
}

func (rest *Rest) Schemes(scheme string) *Rest {
	rest.currentScheme = scheme
	return rest
}

func (rest *Rest) Host(host string) *Rest {
	rest.currentHost = host
	return rest
}

func (rest *Rest) RoutePrefix(prefix string) *Rest {
	rest.currentRoutePrefix = prefix + rest.currentRoutePrefix
	return rest
}

func (rest *Rest) API(version ...string) *Rest {
	rest.currentRoutePrefix = "/api/" + version[0] + rest.currentRoutePrefix
	return rest
}

func (rest *Rest) StrictSlash(value bool) *Rest {
	rest.Mux.StrictSlash(value)
	return rest
}

func (rest *Rest) Route(route string) *Rest {
	rest.currentRoute = route
	return rest
}

func (rest *Rest) Controller(ctrl RestHandler) {
	rest.currentHandler = ctrl
	rest.mapRoute()
	rest.mapControllerHandlers()
}

func (rest *Rest) mapRoute() {
	pathPrefix := rest.currentRoute

	if rest.currentRoutePrefix != "" {
		pathPrefix = rest.currentRoutePrefix + rest.currentRoute
		rest.currentRoutePrefix = ""
	}

	route := rest.Mux.PathPrefix(pathPrefix)

	if rest.currentScheme != "" {
		route = rest.Mux.Schemes(rest.currentScheme)
	}

	if rest.currentHost != "" {
		route = rest.Mux.Host(rest.currentHost)
	}

	rest.SubRouter = route.Subrouter()
}

func (rest *Rest) mapControllerHandlers() {
	ctrl := NewMappedController(rest)

	rest.controllers = append(rest.controllers, ctrl)

	for _, ctrl := range rest.controllers {
		ctrl.Map()
	}
}
