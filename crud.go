package rest

import (
	"net/http"

	"github.com/GoLangWebSDK/crud"
)

type CRUDController[T any] struct {
	Controller
	Repository crud.Repository[T]
}

func NewCRUDController[T any](repo crud.Repository[T]) *CRUDController[T] {
	return &CRUDController[T]{
		Repository: repo,
	}
}

func (ctrl *CRUDController[T]) Create(session *Session) {
	var record T

	err := session.JsonDecode(&record)
	if err != nil {
		session.SetStatus(http.StatusBadRequest)
		return
	}

	newRecord, err := ctrl.Repository.Create(record)
	if err != nil {
		session.SetStatus(http.StatusInternalServerError)
		return
	}

	session.JsonResponse(http.StatusCreated, newRecord)
}

func (ctrl *CRUDController[T]) ReadAll(session *Session) {
	records, err := ctrl.Repository.ReadAll()
	if err != nil {
		session.SetStatus(http.StatusInternalServerError)
		return
	}

	session.JsonResponse(http.StatusOK, records)
}

func (ctrl *CRUDController[T]) Read(session *Session) {
	record, err := ctrl.Repository.Read(session.GetID())
	if err != nil {
		session.JsonResponse(http.StatusNotFound, nil)
		return
	}

	session.JsonResponse(http.StatusOK, record)
}

func (ctrl *CRUDController[T]) Update(session *Session) {
	var record T

	err := session.JsonDecode(&record)
	if err != nil {
		session.SetStatus(http.StatusBadRequest)
		return
	}

	updatedRecord, err := ctrl.Repository.Update(session.GetID(), record)
	if err != nil {
		session.SetStatus(http.StatusInternalServerError)
		return
	}

	session.JsonResponse(http.StatusOK, updatedRecord)
}

func (ctrl *CRUDController[T]) Destroy(session *Session) {
	err := ctrl.Repository.Delete(session.GetID())
	if err != nil {
		session.Response.WriteHeader(http.StatusInternalServerError)
		return
	}

	session.Response.WriteHeader(http.StatusOK)
}

