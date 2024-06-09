package main

import (
	"fmt"
	"github.com/GoLangWebSDK/rest"
)

// Make any struct a controller by embedding rest.Controller
type UsersController struct {
	*rest.Controller
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
