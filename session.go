package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Session struct {
	JsonDecoder *json.Decoder
	Request     *http.Request
	Response    http.ResponseWriter
}

func NewSession(r *http.Request, w http.ResponseWriter) *Session {
	return &Session{
		Request:  r,
		Response: w,
	}
}

func (session *Session) GetParam(key string) string {
	return session.Request.PathValue(key)
}

func (session *Session) GetUUID() string {
	return session.GetParam("uuid")
}

func (session *Session) GetSlug() string {
	return session.GetParam("slug")
}

func (session *Session) GetID(key ...string) uint {
	paramKey := "id"
	if key != nil {
		paramKey = key[0]
	}

	u64, err := strconv.ParseUint(session.GetParam(paramKey), 10, 32)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return uint(u64)
}

func (session *Session) JsonDecode(value interface{}) error {
	session.JsonDecoder = json.NewDecoder(session.Request.Body)
	return session.JsonDecoder.Decode(value)
}

func (session *Session) SetHeader(key, value string) {
	session.Response.Header().Set(key, value)
}

func (session *Session) SetStatus(status int) {
	session.Response.WriteHeader(status)
}

func (session *Session) JsonResponse(status int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.Response.Header().Set("Content-Type", "application/json")
	session.Response.WriteHeader(status)
	session.Response.Write(response)
}
