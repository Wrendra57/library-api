package controller

import (
	"net/http"
	"strconv"

	"github.com/be/perpustakaan/exception"
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
func (c *BookLoanControllerImpl) ReturnBookLoan(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createBookLoanRequest := webrequest.BookLoanCreateRequest{}
	helper.ReadFromRequestBody(request, &createBookLoanRequest)

	returnBook := c.BookLoanService.ReturnBookLoan(request.Context(), createBookLoanRequest)

	webResponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   returnBook,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (c *BookLoanControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	listBookLoanRequest := webrequest.ListALlBookLoanRequest{}
	if request.URL.Query().Get("limit") != "" {
		l, err := strconv.Atoi(request.URL.Query().Get("limit"))
		helper.PanicIfError(err)
		listBookLoanRequest.Limit = l
	}
	if request.URL.Query().Get("limit") != "" {
		o, err := strconv.Atoi(request.URL.Query().Get("offset"))
		helper.PanicIfError(err)
		listBookLoanRequest.Offset = o
	}
	getBookLoan := c.BookLoanService.FindAll(request.Context(), listBookLoanRequest)
	webResponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   getBookLoan,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (c *BookLoanControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "id must be number"})
	}
	bookLoan := c.BookLoanService.FindById(request.Context(), id)
	webResponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   bookLoan,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
func (c *BookLoanControllerImpl) ListByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	list := c.BookLoanService.ListByUserId(request.Context())
	webResponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   list,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
