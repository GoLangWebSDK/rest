REST
---

REST Module allows you to quickly build controller based RESTfull API. It's built on top of `net/http` std lib server multiplexer so no additional dependencies are required.

## Quick Start
Quick example demonstrates how you can simply turn any struct into a rest controller by embbeding the base `rest.Controller` struct. The embbeding implements the `RestHandler` interface with the CRUD methods that are mapped to the request path. To implement custom CRUD handlers just overwrite the interface methods like in the example.

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

Rest module comes with 2 base controller you can use. 

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
Every controller should inherit from this base controller to become the handler for the REST requests. There 2 general ways you can use it:

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
Use custom routing and metods by using 'Run() Controller' method where you can specify your own routing and handlers and overwrite the default methods as well. 

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
The `Run()` method is inspiried with ruby sintra framework that allows you to define your API in more functional way. So you can 
build your API using only `Run() Controller` method and the Get, Post, Put and Delete methods.

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

You can then group and route the requests under a common route to your controller struct.

```go
    router.Route("/users").Controller(NewUsersController(router))
```

Now all the requests with `/users` path prefix will be routed to the `NewUsersController` controller and all the routes in the controller
will append `/users` with default or custom routes defined in the controller.


### Path Prefix
If you need additional prefix to your routes you can use the `RoutePrefix()` method.

```go

    apiRouter := router.RoutePrefix("/api")
    
    // Routed to /api/users
    apiRouter.Route("/users").Controller(NewUsersController(router))

```

In addition to group and version API routes you can use convenience method `API()`.

```go
    
    apiRouter := router.API("/v1")

    // Routed to /api/v1/users 
    apiRouter.Route("/users").Controller(NewUsersController(router))

```

### Middlewares

REST supports middlewares that can be used to intercept and modify the request and response.


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
To handle trailing slashes you can use `router.StrictSlash(false)`, defult is `true`.

```go
    
    router.StrictSlash(false)

    router.Route("/users/").Controller(NewUsersController(router))

```

Active strict slash will strip all your routes that have a trailing slash and redirect them to their non-trailing version. When 
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
REST module comes with `Routes` interface that allows you to load your controller routes and middlewares anywhere in your code base by implementing the interface. This way you can decouple your router config for a more MVC style app development.

```go
    type Routes interface {
	    LoadRoutes(*Router)
	    LoadMiddleware(*Router)
    }
```
In your ie. `app.go` you can do:

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

and then in your main.go you can do:

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

 e## Session 
The REST session is a decorator around the `http.Request` and `http.ResponseWriter` that keeps all the data and some usefull methods for the current request in a single rest session. 

```go

    type Session struct {
	    JsonDecoder *json.Decoder
	    Request     *http.Request
	    Response    http.ResponseWriter
    }
```
Session will help when fetching path values from the request, it is also used to decode incoming request bodies and to repond with JSON responses woth proper status code. 

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

`session.GetID()` - return the session ID as uint by defualt is will search for `{id}` in the URL path if no value is provided, but to fetch id's with custom name you just pass it in as value.

```go

    ctrl.Get("/{id}/posts/{post}", func(session *rest.Session) {
        userID := session.GetID()
        postID := session.GetID("post")
    })

```

Any other parameter can be fetched using `session.GetParam("paramName")`. There are convenience methods the wrapp `GetParam` that return more common params like `slug` or `uuid` returned as strngs


```go

    ctrl.Get("/{uuid}/posts/{slug}", func(session *rest.Session) {
        userUUID := session.GetUUID()
        postSlug := session.GetSlug()
    })

```

To set the status or the header of the reponse use `session.SetStatus` and `session.SetHeader`, `session.JsonResponse` will reponde with JSON, and Content-type will be set to `application/json`. 
