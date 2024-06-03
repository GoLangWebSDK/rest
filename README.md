REST
---

REST Module is built on top [gorilla/mux](https://github.com/gorilla/mux) to allow simple controller routing.

## Quick Start

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/GoLangWebSDK/rest"
)

func main() {
	router := rest.NewRouter()

	apiRouter := router.RoutePrefix("/api")

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

// Mapped to GET /api/users
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
```