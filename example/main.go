package main

import (
	"fmt"
	"net/http"

	"github.com/GoLangWebSDK/rest"
)

func main() {
	router := rest.NewRouter()

	ctrl := rest.NewController(router)

	ctrl.Get("/test/route", func(ctx *rest.Context) {
		fmt.Println("Get /test/route")
	})

	ctrl.Post("/test/route", func(ctx *rest.Context) {
		fmt.Println("Post /test/route")
	})

	apiCtrl := rest.NewController(router)

	apiCtrl.Get("/products", func(ctx *rest.Context) {
		fmt.Println("Get /products")
	})

	apiCtrl.Post("/products", func(ctx *rest.Context) {
		fmt.Println("Post /products")
	})

	// router.Load(app.NewRoutes())

	err := http.ListenAndServe(":8080", router.Mux)
	if err != nil {
		fmt.Println(err)
	}
}
