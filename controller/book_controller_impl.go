package controller

import (
	"fmt"
	"io"
	"net/http"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
	"github.com/be/perpustakaan/service"
	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (c *BookControllerImpl) CreateBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	createRequest := webrequest.BookCreateRequest{}

	createRequest.Title = request.FormValue("title")
	createRequest.Category = request.FormValue("category")
	createRequest.Author = request.FormValue("author")
	createRequest.Publisher = request.FormValue("publisher")
	createRequest.Isbn = request.FormValue("isbn")
	createRequest.Page_count = request.FormValue("page_count")

	createRequest.Stock = request.FormValue("stock")
	createRequest.Publication_year =request.FormValue("publication_year")
	file, _, err := request.FormFile("foto")
	helper.PanicIfError(err)
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	helper.PanicIfError(err)

	createRequest.Foto = fileContents
	createRequest.Rak = request.FormValue("rak")
	createRequest.Column = request.FormValue("column")
	createRequest.Rows = request.FormValue("rows")
	createRequest.Price = request.FormValue("price")
	fmt.Println("control jalan")
	createBook := c.BookService.CreateBook(request.Context(), createRequest)

	webRespone := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   createBook,
	}

	helper.WriteToResponseBody(writer, webRespone)
}