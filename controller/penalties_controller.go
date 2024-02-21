package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PenaltiesController interface {
	PayPenalties(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
