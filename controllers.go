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
	GetPath := "GET " + path

	if ctrl.Router.TrimSlash && path != "/" {
		GetPath = strings.TrimSuffix(GetPath, "/")
	}

	ctrl.createHandler(GetPath, handler)
}

func (ctrl *RestController) Post(path string, handler func(session *Session)) {
	PostPath := "POST " + path

	if ctrl.Router.TrimSlash && path != "/" {
		PostPath = strings.TrimSuffix(PostPath, "/")
	}

	ctrl.createHandler(PostPath, handler)
}

func (ctrl *RestController) Put(path string, handler func(session *Session)) {
	PutPath := "PUT " + path

	if ctrl.Router.TrimSlash && path != "/" {
		PutPath = strings.TrimSuffix(PutPath, "/")
	}

	ctrl.createHandler(PutPath, handler)
}

func (ctrl *RestController) Delete(path string, handler func(session *Session)) {
	DeletePath := "DELETE " + path

	if ctrl.Router.TrimSlash && path != "/" {
		DeletePath = strings.TrimSuffix(DeletePath, "/")
	}

	ctrl.createHandler(DeletePath, handler)
}

func (ctrl *RestController) createHandler(path string, handler func(session *Session)) {
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(NewSession(r, w))
	}

	ctrl.Router.Mux.HandleFunc(path, h)
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
	GetPath := "GET " + ctrl.Router.CurrentRoute.FinalPath + path

	if ctrl.Router.TrimSlash {
		GetPath = strings.TrimSuffix(GetPath, "/")
	}

	ctrl.createHandler(GetPath, handler)
	ctrl.Router.CurrentRoute.mappedPaths = append(ctrl.Router.CurrentRoute.mappedPaths, GetPath)
}

func (ctrl Controller) Post(path string, handler func(session *Session)) {
	PostPath := "POST " + ctrl.Router.CurrentRoute.FinalPath + path

	if ctrl.Router.TrimSlash {
		PostPath = strings.TrimSuffix(PostPath, "/")
	}

	ctrl.createHandler(PostPath, handler)
	ctrl.Router.CurrentRoute.mappedPaths = append(ctrl.Router.CurrentRoute.mappedPaths, PostPath)
}

func (ctrl Controller) Put(path string, handler func(session *Session)) {
	PutPath := "PUT " + ctrl.Router.CurrentRoute.FinalPath + path

	if ctrl.Router.TrimSlash {
		PutPath = strings.TrimSuffix(PutPath, "/")
	}

	ctrl.createHandler(PutPath, handler)
	ctrl.Router.CurrentRoute.mappedPaths = append(ctrl.Router.CurrentRoute.mappedPaths, PutPath)
}

func (ctrl Controller) Delete(path string, handler func(session *Session)) {
	DeletePath := "DELETE " + ctrl.Router.CurrentRoute.FinalPath + path

	if ctrl.Router.TrimSlash {
		DeletePath = strings.TrimSuffix(DeletePath, "/")
	}

	ctrl.createHandler(DeletePath, handler)
	ctrl.Router.CurrentRoute.mappedPaths = append(ctrl.Router.CurrentRoute.mappedPaths, DeletePath)
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
