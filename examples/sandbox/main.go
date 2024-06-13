package main

import (
	"fmt"
	"net/http"

	"github.com/GoLangWebSDK/rest"
)

func main() {
	router := rest.NewRouter()

	router.StrictSlash(true)

	ctrl := rest.NewController(router)

	ctrl.Get("/hello", func(session *rest.Session) {
		fmt.Fprintf(session.Response, "Hello World!")
	})

	ctrl.Get("/check/", func(session *rest.Session) {
		fmt.Fprintf(session.Response, "Route is working!")
	})

	http.ListenAndServe(":8080", router)
}
