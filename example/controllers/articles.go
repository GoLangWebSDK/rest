package controllers

import (
	"fmt"

	"github.com/GoLangWebSDK/rest"
)

type ArticlesController struct {
	rest.Controller
}

func (ctrl *ArticlesController) Run() {
	ctrl.Post("/save", ctrl.Save)
}

func (ctrl *ArticlesController) ReadAll(ctx *rest.Context) {
	fmt.Println("ArticlesController.ReadAll")
}

func (ctrl *ArticlesController) Read(ctx *rest.Context) {
	fmt.Println("ArticlesController.Read")
}

func (ctrl *ArticlesController) Save(ctx *rest.Context) {
	fmt.Println("ArticlesController.Save")
}

func (ctrl *ArticlesController) Destroy(ctx *rest.Context) {
	fmt.Println("ArticlesController.Destroy")
}
