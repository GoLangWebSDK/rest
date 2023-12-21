package controllers

import (
	"fmt"

	"github.com/GoLangWebSDK/rest"
)

type TestController struct {
	rest.Controller	
}

func NewTestController() *TestController {
	return &TestController{}
}

func (ctrl *TestController) Run() {
	ctrl.Get("/test/route", func(ctx *rest.Context) {
		fmt.Println("TestController::Get /test/route")
	})
}

func (ctrl *TestController) ReadAll(ctx *rest.Context) {
	fmt.Println("ApiController::ReadAll")
}