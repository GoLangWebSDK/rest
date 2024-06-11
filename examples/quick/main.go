package main

import (
	"fmt"
	"github.com/GoLangWebSDK/rest"
	"net/http"
)

func main() {
	router := rest.NewRouter()

	router.Mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mux is healthy")
	})

	hello := rest.NewController(router)
	hello.Get("/", func(session *rest.Session) {
		fmt.Fprintf(session.Response, "Controller is healthy")
	})

	router.Route("/users").Controller(NewUsersController(router))

	http.ListenAndServe(":9090", router)

	// fmt.Println("Spinning up server...")
	// router := rest.NewRouter()
	//
	// ctrl := rest.NewController(router)
	//
	// ctrl.Get("/", func(session *rest.Session) {
	// 	fmt.Println("HomeController::Index")
	// 	fmt.Fprintf(session.Response, "Hello World!")
	// })
	//
	// ctrl.Get("/users/{id}", func(session *rest.Session) {
	// 	fmt.Println("UsersController::Read")
	// 	userId := session.GetID()
	//
	// 	var jsonResponse struct {
	// 		Status bool   `json:"status"`
	// 		Msg    string `json:"msg"`
	// 	}
	//
	// 	jsonResponse.Status = true
	// 	jsonResponse.Msg = fmt.Sprintf("User with id %d", userId)
	//
	// 	session.JsonResponse(http.StatusOK, jsonResponse)
	// })
	//
	// ctrl.Post("/users/create", func(session *rest.Session) {
	// 	fmt.Println("UsersController::Create")
	// 	var input struct {
	// 		Username string `json:"username"`
	// 		Password string `json:"password"`
	// 	}
	//
	// 	err := session.JsonDecode(&input)
	// 	if err != nil {
	// 		session.SetStatus(http.StatusBadRequest)
	// 		return
	// 	}
	//
	// 	var responseJson struct {
	// 		Status bool   `json:"status"`
	// 		Msg    string `json:"msg"`
	// 	}
	//
	// 	responseJson.Status = true
	// 	responseJson.Msg = fmt.Sprintf("User %s created", input.Username)
	//
	// 	session.JsonResponse(http.StatusOK, responseJson)
	// })
	//
	// ctrl.Put("/users/{id}/post/{post_id}/update", func(session *rest.Session) {
	// 	fmt.Println("UsersController::Update User Post")
	//
	// 	userId := session.GetID()
	// 	postId := session.GetID("post_id")
	//
	// 	var requestBody struct {
	// 		Title   string `json:"title"`
	// 		Content string `json:"content"`
	// 	}
	//
	// 	err := session.JsonDecode(&requestBody)
	// 	if err != nil {
	// 		session.SetStatus(http.StatusBadRequest)
	// 		return
	// 	}
	//
	// 	var jsonResponse struct {
	// 		Status bool   `json:"status"`
	// 		Msg    string `json:"msg"`
	// 	}
	//
	// 	jsonResponse.Status = true
	// 	jsonResponse.Msg = fmt.Sprintf("User %d updated post %d", userId, postId)
	//
	// 	session.JsonResponse(http.StatusOK, jsonResponse)
	// })
	//
	// ctrl.Delete("/users/{id}", func(session *rest.Session) {
	// 	fmt.Println("UsersController::Destroy User Post")
	//
	// 	userId := session.GetID()
	//
	// 	var jsonResponse struct {
	// 		Status bool   `json:"status"`
	// 		Msg    string `json:"msg"`
	// 	}
	//
	// 	jsonResponse.Status = true
	// 	jsonResponse.Msg = fmt.Sprintf("User %d deleted", userId)
	//
	// 	session.JsonResponse(http.StatusOK, jsonResponse)
	// })
	//
	// fmt.Println("Server started on port 8099")
	// err := http.ListenAndServe(":8099", router)
	// if err != nil {
	// 	fmt.Println("Could not start server, error: ", err)
	// 	return
	// }

}

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
