package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BookController interface {
	CreateBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindBookById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ListBooks(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SearchBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
