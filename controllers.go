package rest

import (
	"net/http"
)

// REST CONTROLLER
type RestController struct {
	Router *http.ServeMux
}

func NewController(router *Router) *RestController {
	return &RestController{
		Router: router.Mux,
	}
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

// (BASE) CONTROLLER
type Controller struct {
	Router *Router
}

var _ RestHandler = &Controller{}

func New(router *Router) *Controller {
	return &Controller{
		Router: router,
	}
}

func (ctrl *Controller) Get(path string, handler HandlerFunc) {
	GetPath := "GET " + ctrl.Router.CurrentRoute.Path + path
	ctrl.createHandler(GetPath, handler)
}

func (ctrl *Controller) Post(path string, handler func(session *Session)) {
	PostPath := "POST " + ctrl.Router.CurrentRoute.Path + path
	ctrl.createHandler(PostPath, handler)
}

func (ctrl *Controller) Put(path string, handler func(session *Session)) {
	PutPath := "PUT " + ctrl.Router.CurrentRoute.Path + path
	ctrl.createHandler(PutPath, handler)
}

func (ctrl *Controller) Delete(path string, handler func(session *Session)) {
	DeletePath := "DELETE " + ctrl.Router.CurrentRoute.Path + path
	ctrl.createHandler(DeletePath, handler)
}

func (ctrl *Controller) createHandler(path string, handler func(session *Session)) {
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(NewSession(r, w))
	}

	ctrl.Router.CurrentRoute.SubRouter.HandleFunc(path, h)
}

func (ctrl *Controller) Run() {}

func (ctrl *Controller) Create(session *Session) {}

func (ctrl *Controller) Read(session *Session) {}

func (ctrl *Controller) ReadAll(session *Session) {}

func (ctrl *Controller) Update(session *Session) {}

func (ctrl *Controller) Destroy(session *Session) {}
