package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerFunc func(ctx *Context)
type Routes interface {
	LoadRoutes()
	LoadMiddlewares()
}

type Rest struct {
	Mux               *mux.Router
	SubRouter         *mux.Router
	CurrentController *RestController
	MapedController   Controller
}

func NewRouter() *Rest {
	router := mux.NewRouter()
	return &Rest{
		Mux:               router,
		SubRouter:         nil,
		CurrentController: &RestController{},
		MapedController:   nil,
	}
}

func (rest *Rest) Init() {

	if rest.MapedController != nil {

		rest.MapedController.Run()

		s := rest.Mux.PathPrefix(rest.CurrentController.Path).Subrouter()

		GetHandler := func(w http.ResponseWriter, r *http.Request) {
			rest.MapedController.Read(NewContext(w, r))
		}

		GetAllHandler := func(w http.ResponseWriter, r *http.Request) {
			rest.MapedController.ReadAll(NewContext(w, r))
		}

		PostHandler := func(w http.ResponseWriter, r *http.Request) {
			rest.MapedController.Create(NewContext(w, r))
		}

		PutHandler := func(w http.ResponseWriter, r *http.Request) {
			rest.MapedController.Update(NewContext(w, r))
		}

		DeleteHandler := func(w http.ResponseWriter, r *http.Request) {
			rest.MapedController.Destroy(NewContext(w, r))
		}

		s.HandleFunc("/", GetAllHandler).Methods("GET")
		s.HandleFunc("/{id}", GetHandler).Methods("GET")
		s.HandleFunc("/", PostHandler).Methods("POST")
		s.HandleFunc("/{id}", PutHandler).Methods("PUT")
		s.HandleFunc("/{id}", DeleteHandler).Methods("DELETE")

		return
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		rest.CurrentController.Handler(NewContext(w, r))
	}

	rest.Mux.HandleFunc(rest.CurrentController.Path, handler).Methods(rest.CurrentController.Methods...)
}

func (rest *Rest) Load(routes Routes) {
	if routes != nil {
		routes.LoadRoutes()
		routes.LoadMiddlewares()
	}
}

func (rest *Rest) MapController(ctrl Controller) *Rest {
	rest.MapedController = ctrl
	rest.SubRouter = rest.Mux.PathPrefix(rest.CurrentController.Path).Subrouter()
	rest.MapedController.Init(rest)
	return rest
}

func (rest *Rest) Controller(handler HandlerFunc) *Rest {
	rest.CurrentController.Handler = handler
	return rest
}

func (rest *Rest) Route(route string) *Rest {
	rest.CurrentController.Path = route
	return rest
}

func (rest *Rest) Methods(methods ...string) *Rest {
	rest.CurrentController.Methods = append(rest.CurrentController.Methods, methods...)
	return rest
}

func (rest *Rest) Listen(port string) error {
	err := http.ListenAndServe(port, rest.Mux)

	if err != nil {
		return err
	}

	return nil
}
