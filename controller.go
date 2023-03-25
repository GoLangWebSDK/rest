package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type controller interface {
	Init(router *Rest)
	Run()
	Create(ctx *Context)
	Read(ctx *Context)
	ReadAll(ctx *Context)
	Update(ctx *Context)
	Destroy(ctx *Context)
}

type Controller struct {
	Router  *Rest
	Path    string
	Handler HandlerFunc
	Methods []string
	Mapped  bool
}

func (ctrl *Controller) Create(ctx *Context)  {}
func (ctrl *Controller) Read(ctx *Context)    {}
func (ctrl *Controller) ReadAll(ctx *Context) {}
func (ctrl *Controller) Update(ctx *Context)  {}
func (ctrl *Controller) Destroy(ctx *Context) {}

func NewController(router *Rest) *Controller {
	return &Controller{
		Router:  router,
		Path:    "",
		Handler: nil,
		Methods: []string{},
		Mapped:  false,
	}
}

func (ctrl *Controller) Init(router *Rest) {
	ctrl.Router = router
	ctrl.Path = router.CurrentController.Path
	ctrl.Methods = router.CurrentController.Methods
	ctrl.Mapped = true
}

func (ctrl *Controller) CreateHandler(path string, handler HandlerFunc) *mux.Route {
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(NewContext(w, r))
	}
	if ctrl.Mapped {
		return ctrl.Router.SubRouter.HandleFunc(path, h)
	}
	return ctrl.Router.Mux.HandleFunc(path, h)
}

func (ctrl *Controller) Get(path string, handler HandlerFunc) {
	ctrl.CreateHandler(path, handler).Methods("GET")
}

func (ctrl *Controller) Post(path string, handler HandlerFunc) {
	ctrl.CreateHandler(path, handler).Methods("POST")
}

func (ctrl *Controller) Put(path string, handler HandlerFunc) {
	ctrl.CreateHandler(path, handler).Methods("PUT")
}

func (ctrl *Controller) Delete(path string, handler HandlerFunc) {
	ctrl.CreateHandler(path, handler).Methods("DELETE")
}
