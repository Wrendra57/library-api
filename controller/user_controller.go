package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Authenticate(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ListAllUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindUserById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
