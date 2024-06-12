package rest

import (
	"fmt"
	"net/http"
)

type RestController struct {
	Router *http.ServeMux
}

func NewController(router *Router) RestController {
	ctrl := RestController{
		Router: router.Mux,
	}
	return ctrl
}

func (ctrl *RestController) Get(path string, handler HandlerFunc) {
	GetPath := "GET " + path
	ctrl.createHandler(GetPath, handler)
}

func (ctrl *RestController) Post(path string, handler func(session *Session)) {
	PostPath := "POST " + path

	ctrl.createHandler(PostPath, handler)
}

func (ctrl *RestController) Put(path string, handler func(session *Session)) {
	PutPath := "PUT " + path
	ctrl.createHandler(PutPath, handler)
}

func (ctrl *RestController) Delete(path string, handler func(session *Session)) {
	DeletePath := "DELETE " + path
	ctrl.createHandler(DeletePath, handler)
}

func (ctrl *RestController) createHandler(path string, handler func(session *Session)) {
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(NewSession(r, w))
	}

	ctrl.Router.HandleFunc(path, h)
}

type Controller struct {
	Router   *Router
	Patterns []string
}

func New(router *Router) Controller {
	return Controller{
		Router: router,
	}
}

func (ctrl Controller) Get(path string, handler HandlerFunc) {
	fmt.Println("Setting custom Get path")
	GetPath := "GET " + ctrl.Router.CurrentRoute.FinalPath + path

	ctrl.createHandler(GetPath, handler)
	ctrl.Router.CurrentRoute.mappedPaths = append(ctrl.Router.CurrentRoute.mappedPaths, GetPath)

	fmt.Println(ctrl.Router.CurrentRoute.mappedPaths)
}

func (ctrl Controller) Post(path string, handler func(session *Session)) {
	PostPath := "POST " + ctrl.Router.CurrentRoute.FinalPath + path
	ctrl.createHandler(PostPath, handler)
}

func (ctrl Controller) Put(path string, handler func(session *Session)) {
	PutPath := "PUT " + ctrl.Router.CurrentRoute.FinalPath + path
	ctrl.createHandler(PutPath, handler)
}

func (ctrl Controller) Delete(path string, handler func(session *Session)) {
	DeletePath := "DELETE " + ctrl.Router.CurrentRoute.FinalPath + path
	ctrl.createHandler(DeletePath, handler)
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
