package main

import (
	"fmt"
	"net/http"

	"github.com/GoLangWebSDK/rest"
)

func main() {
	fmt.Println("Spinning up server...")
	router := rest.NewRouter()
	routes := NewRoutes()

	router.Load(routes)

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}

}
