package controllers

import (
	"fmt"

	"github.com/GoLangWebSDK/rest"
)

type UsersController struct {
	rest.Controller
}

func (ctrl *UsersController) Run() {
	ctrl.Get("/deleted", func(ctx *rest.Context) {
		fmt.Println("UsersController.Get users/deleted")
	})

	ctrl.Get("/active", ctrl.GetActiveUsers)

	ctrl.Post("/save", ctrl.SaveUser)
}

func (ctrl *UsersController) SaveUser(ctx *rest.Context) {
	fmt.Println("UsersController.SaveUser")
}

func (ctrl *UsersController) GetActiveUsers(ctx *rest.Context) {
	fmt.Println("UsersController.GetActiveUsers")
}

func (ctrl *UsersController) ReadAll(ctx *rest.Context) {
	fmt.Println("UsersController.ReadAll")
}

func (ctrl *UsersController) Read(ctx *rest.Context) {
	fmt.Println("UsersController.Read")
}

func (ctrl *UsersController) Create(ctx *rest.Context) {
	fmt.Println("UsersController.Create")
}

func (ctrl *UsersController) Update(ctx *rest.Context) {
	fmt.Println("UsersController.Update")
}

func (ctrl *UsersController) Destroy(ctx *rest.Context) {
	fmt.Println("UsersController.Destroy")
}
