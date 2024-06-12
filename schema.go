package rest

import (
	"net/http"
)

type Routes interface {
	LoadRoutes(*Router)
	LoadMiddleware(*Router)
}

type RestHandler interface {
	Run() Controller
	Create(*Session)
	Read(*Session)
	ReadAll(*Session)
	Update(*Session)
	Destroy(*Session)
}

type HandlerFunc func(session *Session)

type Middleware func(http.Handler) http.Handler
