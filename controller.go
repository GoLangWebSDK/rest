package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	Init(router *Rest)
	Run()
	Create(ctx *Context)
	Read(ctx *Context)
	ReadAll(ctx *Context)
	Update(ctx *Context)
	Destroy(ctx *Context)
}

var _ Controller = (*RestController)(nil)

type RestController struct {
	Router  *Rest
	Path    string
	Handler HandlerFunc
	Methods []string
	Mapped  bool
}

func NewController(router *Rest) *RestController {
	return &RestController{
		Router:  router,
		Path:    "",
		Handler: nil,
		Methods: []string{},
		Mapped:  false,
	}
}

func (ctrl *RestController) Run()  								{}
func (ctrl *RestController) Create(ctx *Context)  {}
func (ctrl *RestController) Read(ctx *Context)    {}
func (ctrl *RestController) ReadAll(ctx *Context) {}
func (ctrl *RestController) Update(ctx *Context)  {}
func (ctrl *RestController) Destroy(ctx *Context) {}

func (ctrl *RestController) Init(router *Rest) {
	ctrl.Router = router
	ctrl.Path = router.CurrentController.Path
	ctrl.Methods = router.CurrentController.Methods
	ctrl.Mapped = true
}

func (ctrl *RestController) Get(path string, handler HandlerFunc) {
	ctrl.createHandler(path, handler).Methods("GET")
}

func (ctrl *RestController) Post(path string, handler HandlerFunc) {
	ctrl.createHandler(path, handler).Methods("POST")
}

func (ctrl *RestController) Put(path string, handler HandlerFunc) {
	ctrl.createHandler(path, handler).Methods("PUT")
}

func (ctrl *RestController) Delete(path string, handler HandlerFunc) {
	ctrl.createHandler(path, handler).Methods("DELETE")
}

func (ctrl *RestController) createHandler(path string, handler HandlerFunc) *mux.Route {
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(NewContext(w, r))
	}
	if ctrl.Mapped {
		return ctrl.Router.SubRouter.HandleFunc(path, h)
	}
	return ctrl.Router.Mux.HandleFunc(path, h)
}