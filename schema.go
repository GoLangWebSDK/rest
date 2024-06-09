package rest

import "net/http"

type Routes interface {
	LoadRoutes(router *Rest)
	LoadMiddlewares(router *Rest)
}

type RestHandler interface {
	Run()
	Create(session *Session)
	Read(session *Session)
	ReadAll(session *Session)
	Update(session *Session)
	Destroy(session *Session)
}

type HandlerFunc func(session *Session)

type Middleware func(handler http.Handler) http.Handler
