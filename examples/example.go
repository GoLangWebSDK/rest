package rest

import (
	"fmt"
	"net/http"

	"github.com/GoLangWebSDK/rest"
)

func main() {
	router := rest.NewRouter()

	apiRouter := router.PathPrefix("/api")

	apiRouter.Route("/users").Controller(NewUsersController())

	err := http.ListenAndServe(":8080", router.Mux)
	if err != nil {
		fmt.Println(err)
	}
}


type UsersController struct {
	rest.Controller
}

func NewUsersController() *UsersController {
	return &UsersController{}
}