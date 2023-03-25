package rest

import (
	"gitlab.sintezis.co/sintezis/sdk/web/dmp"

	"gorm.io/gorm"
)

type CRUDController[T any] struct {
	Controller
	Repository *dmp.Repository[T]
	Model      T
}

func (ctrl *CRUDController[T]) Run() {
	ctrl.Repository = dmp.NewRepository(ctrl.Model)
}

func (ctrl *CRUDController[T]) Create(ctx *Context) {
	var input T
	err := ctx.JsonDecode(&input)

	if err != nil {
		ctx.Response.WriteHeader(400)
		return
	}

	err = ctrl.Repository.Add(input)

	if err != nil {
		ctx.Response.WriteHeader(500)
		return
	}

	ctx.Response.WriteHeader(201)
}

func (ctrl *CRUDController[T]) Read(ctx *Context) {
	ctx.SetContentType("application/json")
	ID := ctx.GetID()
	result, err := ctrl.Repository.Get(ID)

	if err != nil {
		ctx.JsonResponse(404, "Not found.")
		return
	}

	ctx.JsonResponse(200, result)
}

func (ctrl *CRUDController[T]) ReadAll(ctx *Context) {
	ctx.SetContentType("application/json")
	results, err := ctrl.Repository.GetAll()

	if err != nil {
		ctx.JsonResponse(500, "Internal server error.")
		return
	}

	ctx.JsonResponse(200, results)
}

func (ctrl *CRUDController[T]) Update(ctx *Context) {
	ID := ctx.GetID()
	var input T

	err := ctx.JsonDecode(&input)

	if err != nil {
		ctx.Response.WriteHeader(400)
		return
	}

	err = ctrl.Repository.Update(ID, input)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JsonResponse(404, "Not found.")
			return
		}
		ctx.JsonResponse(500, "Internal server error.")
		return
	}

	ctx.Response.WriteHeader(200)
}

func (ctrl *CRUDController[T]) Destroy(ctx *Context) {
	ID := ctx.GetID()

	err := ctrl.Repository.Delete(ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JsonResponse(404, "Not found.")
			return
		}
		ctx.Response.WriteHeader(500)
		return
	}

	ctx.Response.WriteHeader(200)
}
