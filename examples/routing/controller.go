package main

import (
	"fmt"
	"net/http"

	"github.com/GoLangWebSDK/rest"
)

type PostsController struct {
	rest.Controller
}

func NewPostsController(router *rest.Router) *PostsController {
	return &PostsController{
		Controller: rest.New(router),
	}
}

func (ctrl *PostsController) Run() {

	ctrl.Post("/", func(session *rest.Session) {
		fmt.Println("Creating Post...")
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
		jsonResponse.Msg = fmt.Sprintf("Created post %s", requestBody.Title)
		session.JsonResponse(http.StatusOK, jsonResponse)
	})

	ctrl.Get("/{slug}", func(session *rest.Session) {

		postSlug := session.GetID()

		var jsonResponse struct {
			Status bool   `json:"status"`
			Msg    string `json:"msg"`
		}

		jsonResponse.Status = true
		jsonResponse.Msg = fmt.Sprintf("Post with slug %d", postSlug)

		session.JsonResponse(http.StatusOK, jsonResponse)
	})

	ctrl.Get("/filter/{key}/{value}", func(session *rest.Session) {

		key := session.GetParam("key")
		value := session.GetParam("value")

		var jsonResponse struct {
			Status bool   `json:"status"`
			Msg    string `json:"msg"`
		}

		jsonResponse.Status = true
		jsonResponse.Msg = fmt.Sprintf("Post filtered by %s with value %s", key, value)

		session.JsonResponse(http.StatusOK, jsonResponse)
	})

	ctrl.Put("/{slug}", func(session *rest.Session) {
		postSlug := session.GetSlug()

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
		jsonResponse.Msg = fmt.Sprintf("Updated post %s, with content: %s", postSlug, requestBody.Content)

		session.JsonResponse(http.StatusOK, jsonResponse)
	})

	ctrl.Delete("/{slug}", func(session *rest.Session) {

		postSlug := session.GetSlug()

		var jsonResponse struct {
			Status bool   `json:"status"`
			Msg    string `json:"msg"`
		}

		jsonResponse.Status = true
		jsonResponse.Msg = fmt.Sprintf("Deleted post %s", postSlug)

		session.JsonResponse(http.StatusOK, jsonResponse)
	})
}
