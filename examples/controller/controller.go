package main

import (
	"fmt"
	"github.com/GoLangWebSDK/rest"
)

// Make any struct a controller by embedding rest.Controller
type UsersController struct {
	rest.Controller
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (ctrl *UsersController) Run() {

}

// Mapped to POST /api/users
func (ctrl *UsersController) Create(ctx *rest.Context) {
	fmt.Println("UsersController::Create")
}

// Mapped to GET /api/users/{id}
func (ctrl *UsersController) Read(ctx *rest.Context) {
	fmt.Println("UsersController::Read")
}

// Mapped to GET /api/users
func (ctrl *UsersController) ReadAll(ctx *rest.Context) {
	fmt.Println("UsersController::ReadAll")
}

// Mapped to PUT /api/users/{id}
func (ctrl *UsersController) Update(ctx *rest.Context) {
	fmt.Println("UsersController::Update")
}

// Mapped to DELETE /api/users/{id}
func (ctrl *UsersController) Destroy(ctx *rest.Context) {
	fmt.Println("UsersController::Destroy")
}
