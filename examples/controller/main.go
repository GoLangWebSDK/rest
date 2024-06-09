package main

import (
	"fmt"
	"net/http"

	"github.com/GoLangWebSDK/rest"
)

func main() {
	router := rest.NewRouter()

	apiRouter := router.RoutePrefix("/api")

	apiRouter.Route("/users").Controller(NewUsersController(router))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
