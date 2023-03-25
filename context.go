package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Context struct {
	Request     *http.Request
	Response    http.ResponseWriter
	Params      map[string]string
	Query       map[string]string
	JsonDecoder *json.Decoder
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request:  r,
		Response: w,
	}
}

func (ctx *Context) GetParams() map[string]string {
	ctx.Params = mux.Vars(ctx.Request)
	return ctx.Params
}

func (ctx *Context) GetParam(key string) string {
	ctx.Params = mux.Vars(ctx.Request)
	return ctx.Params[key]
}

func (ctx *Context) GetID() uint32 {
	u64, err := strconv.ParseUint(ctx.GetParam("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(u64)
}

func (ctx *Context) SetHeader(key, value string) {
	ctx.Response.Header().Set(key, value)
}

func (ctx *Context) SetContentType(contentType string) {
	ctx.Response.Header().Set("Content-Type", contentType)
}

func (ctx *Context) SetStatus(status int) {
	ctx.Response.WriteHeader(status)
}

func (ctx *Context) JsonDecode(value any) error {
	ctx.JsonDecoder = json.NewDecoder(ctx.Request.Body)
	return ctx.JsonDecoder.Decode(value)
}

func (ctx *Context) JsonResponse(status int, data interface{}) error {
	jsonResponse, jsonError := json.Marshal(data)

	if jsonError != nil {
		return jsonError
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	ctx.Response.WriteHeader(status)
	ctx.Response.Write(jsonResponse)

	return nil
}

func (ctx *Context) HtmlResponse(status int, body string) {
	ctx.Response.Header().Set("Content-Type", "text/html")
	ctx.Response.WriteHeader(status)
	ctx.Response.Write([]byte(body))
}
