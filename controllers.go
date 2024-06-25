package rest

import (
	"net/http"
	"strings"
)

type RestController struct {
	Router *Router
}

func NewController(router *Router) RestController {
	ctrl := RestController{
		Router: router,
	}
	return ctrl
}

func (ctrl *RestController) Get(path string, handler HandlerFunc) {
	GetPath := ctrl.buildPath("GET", path)
	ctrl.createHandler(GetPath, handler)
}

func (ctrl *RestController) Post(path string, handler func(session *Session)) {
	PostPath := ctrl.buildPath("POST", path)
	ctrl.createHandler(PostPath, handler)
}

func (ctrl *RestController) Put(path string, handler func(session *Session)) {
	PutPath := ctrl.buildPath("PUT", path)
	ctrl.createHandler(PutPath, handler)
}

func (ctrl *RestController) Delete(path string, handler func(session *Session)) {
	DeletePath := ctrl.buildPath("DELETE", path)
	ctrl.createHandler(DeletePath, handler)
}

func (ctrl *RestController) createHandler(path string, handler func(session *Session)) {
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(NewSession(r, w))
	}

	ctrl.Router.Mux.HandleFunc(path, h)
}

func (ctrl *RestController) buildPath(method string, path string) string {
	route := method + " " + path

	if ctrl.Router.TrimSlash && path != "/" {
		route = strings.TrimSuffix(route, "/")
	}

	return route
}

type Controller struct {
	Router *Router
}

func New(router *Router) Controller {
	return Controller{
		Router: router,
	}
}

func (ctrl Controller) Get(path string, handler HandlerFunc) {
	GetPath := ctrl.buildPath("GET", path)
	ctrl.createHandler(GetPath, handler)
}

func (ctrl Controller) Post(path string, handler func(session *Session)) {
	PostPath := ctrl.buildPath("POST", path)
	ctrl.createHandler(PostPath, handler)
}

func (ctrl Controller) Put(path string, handler func(session *Session)) {
	PutPath := ctrl.buildPath("PUT", path)
	ctrl.createHandler(PutPath, handler)
}

func (ctrl Controller) Delete(path string, handler func(session *Session)) {
	DeletePath := ctrl.buildPath("DELETE", path)
	ctrl.createHandler(DeletePath, handler)
}

func (ctrl Controller) buildPath(method string, path string) string {
	route := method + " " + ctrl.Router.CurrentRoute.FinalPath + path

	if ctrl.Router.TrimSlash {
		route = strings.TrimSuffix(route, "/")
	}

	ctrl.Router.CurrentRoute.mappedPaths = append(ctrl.Router.CurrentRoute.mappedPaths, route)

	return route
}

func (ctrl Controller) createHandler(path string, handler func(session *Session)) {
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(NewSession(r, w))
	}

	ctrl.Router.CurrentRoute.SubRouter.HandleFunc(path, h)
}

func (ctrl Controller) Run() Controller {
	return ctrl
}

func (ctrl Controller) Create(session *Session) {}

func (ctrl Controller) Read(session *Session) {}

func (ctrl Controller) ReadAll(session *Session) {}

func (ctrl Controller) Update(session *Session) {}

func (ctrl Controller) Destroy(session *Session) {}
