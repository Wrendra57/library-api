package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BookLoanController interface {
	CreateBookLoan(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
