package main

import (
	"github.com/GoLangWebSDK/rest"

	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Spinning up server...")
	router := rest.NewRouter()

	router.Use(AuthMiddleware)

	ctrl := rest.NewController(router)
	ctrl.Get("/", func(session *rest.Session) {
		fmt.Println("HomeController::Index")
		fmt.Fprintf(session.Response, "Hello World!")
	})

	ctrl.Get("/users/{id}", func(session *rest.Session) {
		fmt.Println("UsersController::Read")
		userId := session.GetID()

		var jsonResponse struct {
			Status bool   `json:"status"`
			Msg    string `json:"msg"`
		}

		jsonResponse.Status = true
		jsonResponse.Msg = fmt.Sprintf("User with id %d", userId)

		session.JsonResponse(http.StatusOK, jsonResponse)
	})

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

	ctrl.Put("/users/{id}/post/{post_id}/update", func(session *rest.Session) {
		fmt.Println("UsersController::Update User Post")

		userId := session.GetID()
		postId := session.GetID("post_id")

		var requestBody struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		err := session.JsonDecode(&requestBody)
		if err != nil {
			session.SetStatus(http.StatusBadRequest)
			return
		}

		if requestBody.Content == "" || requestBody.Title == "" {
			session.SetStatus(http.StatusBadRequest)
			return
		}

		var jsonResponse struct {
			Status bool   `json:"status"`
			Msg    string `json:"msg"`
		}

		jsonResponse.Status = true
		jsonResponse.Msg = fmt.Sprintf("User %d updated post %d", userId, postId)

		session.JsonResponse(http.StatusOK, jsonResponse)
	})

	ctrl.Delete("/users/{id}", func(session *rest.Session) {
		fmt.Println("UsersController::Destroy User Post")

		userId := session.GetID()

		var jsonResponse struct {
			Status bool   `json:"status"`
			Msg    string `json:"msg"`
		}

		jsonResponse.Status = true
		jsonResponse.Msg = fmt.Sprintf("User %d deleted", userId)

		session.JsonResponse(http.StatusOK, jsonResponse)
	})

	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Could not start server, error: ", err)
		return
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running Auth Middleware for Request: ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
