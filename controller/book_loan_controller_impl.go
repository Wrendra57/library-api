package controller

import (
	"net/http"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
	"github.com/be/perpustakaan/service"
	"github.com/julienschmidt/httprouter"
)

type BookLoanControllerImpl struct {
	BookLoanService service.BookLoanService
}

func NewBookLoanController(bookLoanService service.BookLoanService) BookLoanController {
	return &BookLoanControllerImpl{
		BookLoanService: bookLoanService,
	}
}

func (c *BookLoanControllerImpl) CreateBookLoan(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createBookLoanRequest := webrequest.BookLoanCreateRequest{}
	helper.ReadFromRequestBody(request, &createBookLoanRequest)

	createBookLoan := c.BookLoanService.CreateBookLoan(request.Context(), createBookLoanRequest)

	webResponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   createBookLoan,
	}
	helper.WriteToResponseBody(writer, webResponse)

}
