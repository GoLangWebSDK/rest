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

// Make any struct a controller by embedding rest.Controller
type UsersController struct {
	rest.Controller
}

func NewUsersController(router *rest.Router) *UsersController {
	return &UsersController{
		Controller: rest.New(router),
	}
}

// Mapped to POST /api/users
func (ctrl *UsersController) Create(session *rest.Session) {
	fmt.Println("UsersController::Create")
}

// Mapped to GET /api/users/{id}
func (ctrl *UsersController) Read(session *rest.Session) {
	fmt.Println("UsersController::Read")
}

// Mapped to GET /api/users
func (ctrl *UsersController) ReadAll(session *rest.Session) {
	fmt.Println("UsersController::ReadAll")
}

// Mapped to PUT /api/users/{id}
func (ctrl *UsersController) Update(session *rest.Session) {
	fmt.Println("UsersController::Update")
}

// Mapped to DELETE /api/users/{id}
func (ctrl *UsersController) Destroy(session *rest.Session) {
	fmt.Println("UsersController::Destroy")
}
