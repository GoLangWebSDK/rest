REST
---

REST Module allows you to quickly build controller based RESTfull API. It's built on top of `net/http` std lib  multiplexer so no additional dependencies are required.

## Quick Start
Quick example demonstrates how you can simply turn any struct into a rest controller by embbeding the base `rest.Controller` struct. The embbeding implements the `RestHandler` interface with the CRUD methods that are mapped to the request paths. To implement custom CRUD handlers just overwrite the interface methods like in the example.

Checkout the [examples](https://github.com/GoLangWebSDK/rest/tree/master/examples) for more working examples.

### Install
```shell
go get github.com/GoLangWebSDK/rest
```

### Usage
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

func NewUsersController(router *rest.Router) *UsersController {
	return &UsersController{
		Controller: rest.New(router)
	}
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
## Controllers

Rest module comes with 2 base controllers you can use. 

- `Controller` - This is the base controller that all other controllers inherit from.
- `RestController` - Helps you to express REST patterns as methods on a controller.

### RestController

```go

func main() {
	router := rest.NewRouter()

	router.StrictSlash(true)

	ctrl := rest.NewController(router)

	ctrl.Get("/hello", func(session *rest.Session) {
		fmt.Fprintf(session.Response, "Hello World!")
	})

	ctrl.Get("/check/", func(session *rest.Session) {
		fmt.Fprintf(session.Response, "Route is healthy!")
	})

	ctrl.Post("/create", func(session *rest.Session) {
		fmt.Fprintf(session.Response, "Post route is working!")		
	})

	ctrl.Put("/update", func(session *rest.Session) {
		fmt.Fprintf(session.Response, "Put route is working!")
	})

	ctrl.Delete("/delete", func(session *rest.Session) {
		fmt.Fprintf(session.Response, "Delete route is working!")
	})

	http.ListenAndServe(":8080", router)
}
```

### (Base) Controller

Base controller is the defualt implementation of the `RestHandler` interface.

```go
type RestHandler interface {
	Run() Controller
	Create(*Session)
	Read(*Session)
	ReadAll(*Session)
	Update(*Session)
	Destroy(*Session)
}
```
Every controller should embbed `rest.Controller` to become the handler for the REST requests. There 2 general ways you can use it:

Use the default methods (Create, Read, ReadAll, Update, Destroy)

```go

type UsersController struct {
	rest.Controller
}

func NewUsersController(router *rest.Router) *UsersController {
	return &UsersController{
		Controller: rest.New(router),
	}
}

// Executes before any of the default methods
func (ctrl *UsersController) Run() Controller {
	return ctrl.Controller
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
Use custom routing and methods by using 'Run() Controller' to specify custom routes, handlers and overwrite the default handlers. 

```go

type UsersController struct {
	rest.Controller
}

func NewUsersController() *UsersController {
	return &UsersController{
		Controller: rest.New(router),
	}
}

func (ctrl *UsersController) Run() Controller {
	ctrl.Get("/", GetUser)
	
	return ctrl.Controller
}

func (ctrl *UsersController) GetUser(ctx *rest.Context) {
	fmt.Println("UsersController::GetUser")	
}
```
The `Run()` method is inspired by ruby sinatra framework that allows you to define your routing in a more DSL way.
Build your API using only `Run() Controller` method and the Get, Post, Put and Delete handlers.

```go

type UsersController struct {
	rest.Controller
}

func NewUsersController() *UsersController {
	return &UsersController{
		Controller: rest.New(router),
	}
}

func (ctrl *UsersController) Run() Controller {
	ctrl.Get("/", func(session *rest.Session) {
		fmt.Println("UsersController::GetUsers")
	})

	ctrl.Get("/{id}", func(session *rest.Session) {
		fmt.Println("UsersController::GetUser")
	})

	ctrl.Post("/", func(session *rest.Session) {
		fmt.Println("UsersController::CerateUsers")
	})

	ctrl.Put("/{id}", func(session *rest.Session) {
		fmt.Println("UsersController::UpdateUser")
	})

	ctrl.Delete("/{id}", func(session *rest.Session) {
		fmt.Println("UsersController::DeleteUser")
	})

	return ctrl.Controller
}
```
## Routing

Group and route the requests under a common path to your controller struct.

```go
    router.Route("/users").Controller(NewUsersController(router))
```

Now all the requests with `/users` path prefix will be routed to the `UsersController` controller and all the routes in the controller
will append `/users` with default or custom routes defined in the controller.


### Path Prefix
If you need additional prefix for the routes use the `RoutePrefix()` method.

```go

    apiRouter := router.RoutePrefix("/api")
    
    // Routed to /api/users
    apiRouter.Route("/users").Controller(NewUsersController(router))

```

To group and verions API routes:

```go
    
    apiRouter := router.API("/v1")

    // Routed to /api/v1/users 
    apiRouter.Route("/users").Controller(NewUsersController(router))

```

### Middlewares

REST supports middleware that can be used to intercept and modify the request and response.


```go
    
    router.Use(Logger)

    router.Route("/users").Controller(NewUsersController(router))

    func Logger(next http.Handler) http.Handler {
	    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		    fmt.Println("Running Logger Middleware for Request: ", r.URL.Path)
		    next.ServeHTTP(w, r)
	    })
    }

```

### Strict Slash
To handle trailing slashes use `router.StrictSlash(false)`, defult is `true`.

```go
    
    router.StrictSlash(false)

    router.Route("/users/").Controller(NewUsersController(router))

```

Strict slash set to `true` will strip all your routes with a trailing slash and redirect them to their non-trailing version. When 
set to `false` the routes will be left intact and you have to handle the trailing slash manually.

Strict slash is true: 

`/users/` route will be mapped as `/users` and incoming requests to `/users/` will be redirected to `/users`.

Strict slash is false:

You need to manually handle the trailing slash.

```go
    
    router.StrictSlash(false)
    
    router.Route("/users").Controller(NewUsersController(router))
    router.Route("/users/").Controller(NewUsersController(router))

```

### Routes 
REST module comes with `Routes` interface that allows loasing controller routes and middlewares anywhere in the codebase by implementing the interface. 

```go
    type Routes interface {
	    LoadRoutes(*Router)
	    LoadMiddleware(*Router)
    }
```
Ie. `app.go`:

```go


import (
	"fmt"

	"github.com/GoLangWebSDK/rest"
)

type Routes struct{}

func NewRoutes() *Routes {
	return &Routes{}
}

var _ rest.Routes = &Routes{}

func (routes *Routes) LoadRoutes(router *rest.Router) {
	// router.API() will add a prefix /api to all subsequent routes.
	// It also allows you to set version, for example router.API("/v1"),
	// in which case the final path prefix will be /api/v1
	apiRouter := router.API()

	apiRouter.Route("/posts").Controller(NewPostsController(router))

	fmt.Println("Current route final path: ", apiRouter.CurrentRoute.FinalPath)

}

func (routes *Routes) LoadMiddleware(router *rest.Router) {
	router.Use(
		Logger,
		Auth,
	)
}
```

`main.go`

```go


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

```

### Session 
The REST session is a decorator around the `http.Request` and `http.ResponseWriter` that keeps all the data and some usefull methods for the current request in a single rest session. 

```go

    type Session struct {
	    JsonDecoder *json.Decoder
	    Request     *http.Request
	    Response    http.ResponseWriter
    }
```
Session allows fetching path values from the request, it is also used to decode incoming request bodies and to format JSON responses with proper status code. 

```go
	ctrl.Post("/users/create", func(session *rest.Session) {
		fmt.Println("UsersController::Create")
		var requestBody struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		err := session.JsonDecode(&requestBody)
		if err != nil {
			session.SetStatus(http.StatusBadRequest)
			return
		}

		if requestBody.Password == "" || requestBody.Username == "" {
			session.SetStatus(http.StatusBadRequest)
			return
		}

		var responseJson struct {
			Status bool   `json:"status"`
			Msg    string `json:"msg"`
		}

		responseJson.Status = true
		responseJson.Msg = fmt.Sprintf("User %s created", requestBody.Username)

		session.JsonResponse(http.StatusOK, responseJson)
	})
```

`session.GetID()` - return the session ID as `uint`, will search for `{id}` in the URL path if no value is provided.

```go

    ctrl.Get("/{id}/posts/{post}", func(session *rest.Session) {
        userID := session.GetID()
        postID := session.GetID("post")
    })

```
- set the reponse status - `session.SetStatus` 
- set response header  - `session.SetHeader`, 
- reponde with JSON, and set Content-type `application/json` - `session.JsonResponse` 
- fetch path parameters - `session.GetParam("paramName")`

Fetch common params like `slug` or `uuid`:

```go

    ctrl.Get("/{uuid}/posts/{slug}", func(session *rest.Session) {
        userUUID := session.GetUUID()
        postSlug := session.GetSlug()
    })

```

