package main

import (
	"fmt"
	"net/http"

	"github.com/GoLangWebSDK/rest"
)

func main() {
	fmt.Println("Spinning up server...")
	router := rest.NewRouter()

	apiRouter := router.RoutePrefix("/api")

	apiRouter.Route("/users").Controller(NewUsersController(router))

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

// Make any struct a controller by embedding rest.Controller
type UsersController struct {
	rest.Controller
}

func (ctrl *UsersController) Run() rest.Controller {

	ctrl.Get("/", func(session *rest.Session) {
		fmt.Println("UsersController::Run::ReadAll")

		var jsonResponse struct {
			Status bool   `json:"status"`
			Msg    string `json:"msg"`
		}

		jsonResponse.Status = true
		jsonResponse.Msg = "All users"
		session.JsonResponse(http.StatusOK, jsonResponse)
	})

	return ctrl.Controller
}

func NewUsersController(router *rest.Router) *UsersController {
	return &UsersController{
		Controller: rest.New(router),
	}
}

// Mapped to POST /api/users
func (ctrl *UsersController) Create(session *rest.Session) {
	fmt.Println("UsersController::Create User")

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
}

// Mapped to GET /api/users/{id}
func (ctrl *UsersController) Read(session *rest.Session) {
	fmt.Println("UsersController::Read")

	userId := session.GetID()

	var jsonResponse struct {
		Status bool   `json:"status"`
		Msg    string `json:"msg"`
	}

	jsonResponse.Status = true
	jsonResponse.Msg = fmt.Sprintf("User with id %d", userId)
	session.JsonResponse(http.StatusOK, jsonResponse)
}

// Mapped to GET /api/users
func (ctrl *UsersController) ReadAll(session *rest.Session) {
	fmt.Println("UsersController::ReadAll")

	var jsonResponse struct {
		Status bool   `json:"status"`
		Msg    string `json:"msg"`
	}

	jsonResponse.Status = true
	jsonResponse.Msg = "All users"
	session.JsonResponse(http.StatusOK, jsonResponse)
}

// Mapped to PUT /api/users/{id}
func (ctrl *UsersController) Update(session *rest.Session) {
	fmt.Println("UsersController::Update")

	userId := session.GetID()

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
	responseJson.Msg = fmt.Sprintf("User %s, with id %d updated", requestBody.Username, userId)
	session.JsonResponse(http.StatusOK, responseJson)
}

// Mapped to DELETE /api/users/{id}
func (ctrl *UsersController) Destroy(session *rest.Session) {
	fmt.Println("UsersController::Destroy")

	userId := session.GetID()

	var responseJson struct {
		Status bool   `json:"status"`
		Msg    string `json:"msg"`
	}

	responseJson.Status = true
	responseJson.Msg = fmt.Sprintf("User %d deleted", userId)
	session.JsonResponse(http.StatusOK, responseJson)
}
