package main

import (
	"github.com/GoLangWebSDK/rest"
)

type Routes struct{}

func NewRoutes() *Routes {
	return &Routes{}
}

var _ rest.Routes = &Routes{}

func (routes *Routes) LoadRoutes(router *rest.Router) {
	// router.API() will add a prefix /api to all subsequent routes.
	// It also allows you to set version, for example router.API("/v1"),
	// in which case the final path prefix will be /api/v1
	apiRouter := router.API()

	apiRouter.Route("/posts").Controller(NewPostsController(router))
}

func (routes *Routes) LoadMiddleware(router *rest.Router) {
	router.Use(
		Logger,
		Auth,
	)
}
