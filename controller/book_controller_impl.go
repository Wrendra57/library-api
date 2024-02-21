package controller

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/helper/konversi"
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
	createRequest.Publication_year = request.FormValue("publication_year")
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
	// fmt.Println("control jalan")
	createBook := c.BookService.CreateBook(request.Context(), createRequest)

	webRespone := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   createBook,
	}

	helper.WriteToResponseBody(writer, webRespone)
}

func (c *BookControllerImpl) FindBookById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "id must be number"})
	}
	// fmt.Println("s")
	book := c.BookService.FindBookById(request.Context(), id)

	webresponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   book,
	}
	helper.WriteToResponseBody(writer, webresponse)
}

func (c *BookControllerImpl) ListBooks(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	requestGet := webrequest.FindAllRequest{}

	requestGet.Limit = request.URL.Query().Get("limit")
	requestGet.Offset = request.URL.Query().Get("offset")

	book := c.BookService.ListBook(request.Context(), requestGet)

	webresponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   book,
	}
	helper.WriteToResponseBody(writer, webresponse)
}

func (c *BookControllerImpl) SearchBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	search := request.URL.Query().Get("search")

	requestLimit := webrequest.FindAllRequest{}
	requestLimit.Limit = request.URL.Query().Get("limit")
	requestLimit.Offset = request.URL.Query().Get("offset")

	book := c.BookService.SearchBook(request.Context(), search, requestLimit)

	webresponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   book,
	}
	helper.WriteToResponseBody(writer, webresponse)
}

func (c *BookControllerImpl) UpdateBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	id := konversi.StrToInt(params.ByName("id"), "book id")
	// helper.PanicIfError(err)

	// fmt.Println("s")

	updateRequest := webrequest.UpdateBookRequest{}

	for key, values := range request.Form {
		if len(values) > 0 {
			switch key {
			case "title":
				updateRequest.Title = values[0]
			case "category":
				updateRequest.Category = values[0]
			case "author":
				updateRequest.Author = values[0]
			case "publisher":
				updateRequest.Publisher = values[0]
			case "isbn":
				updateRequest.Isbn = values[0]
			case "page_count":
				updateRequest.Page_count = values[0]
			case "stock":
				updateRequest.Stock = values[0]
			case "publication_year":
				updateRequest.Publication_year = values[0]
			case "rak":
				updateRequest.Rak = values[0]
			case "column":
				updateRequest.Column = values[0]
			case "rows":
				updateRequest.Rows = values[0]
			case "price":
				updateRequest.Price = values[0]
			}
		}
	}
	// if request.FormFile("")
	file, _, err := request.FormFile("foto")
	// fmt.Println(file)
	if err != nil {
		fmt.Println("gada file")
	}
	// defer file.Close()

	var foto []byte

	if file != nil {
		fileContents, err := io.ReadAll(file)
		helper.PanicIfError(err)
		foto = fileContents
		defer file.Close()
	}
	updateRequest.Foto = foto

	update := c.BookService.UpdateBook(request.Context(), updateRequest, id)

	webresponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   update,
	}
	helper.WriteToResponseBody(writer, webresponse)
}

func (c *BookControllerImpl) DeleteBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "id must be number"})
	}
	deleteBook := c.BookService.DeleteBook(request.Context(), id)

	webresponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   deleteBook,
	}
	helper.WriteToResponseBody(writer, webresponse)
}
