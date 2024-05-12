package main

import (
	"fmt"
	"github.com/GoLangWebSDK/rest"
	"net/http"
)

func main() {
	router := rest.NewRouter()

	ctrl := rest.NewController(router)

	ctrl.Get("/users/{id}", func(ctx *rest.Context) {
		fmt.Println("UsersController::Read")
	})

	ctrl.Post("/save/users", func(ctx *rest.Context) {
		fmt.Println("UsersController::Create")
	})

	ctrl.Post("/save/users/{id}/post", func(ctx *rest.Context) {
		fmt.Println("UsersController::Update")
	})

	err := http.ListenAndServe(":8080", router.Mux)
	if err != nil {
		fmt.Println("Could not start server, error: ", err)
		return
	}
}
