package controller

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
	"github.com/be/perpustakaan/service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)
	registerRequest := webrequest.UserCreateRequest{}
	registerRequest.Name = request.FormValue("name")
	registerRequest.Email = request.FormValue("email")
	registerRequest.Password = request.FormValue("password")
	registerRequest.Gender = request.FormValue("gender")
	registerRequest.Telp = request.FormValue("telp")

	// parsing date
	layout := "2006-01-02"
	if request.FormValue("birthdate") == "" {
		panic(exception.CustomEror{Code: 400, Error: "Birthdate is required"})
	}
	parsedTime, err := time.Parse(layout, request.FormValue("birthdate"))
	if err != nil {

		panic(exception.CustomEror{Code: 400, Error: "Birthdate must be format YYYY-MM-DD"})
	}

	registerRequest.Birthdate = parsedTime
	registerRequest.Address = request.FormValue("address")

	file, _, err := request.FormFile("foto")
	if err != nil {
		if err.Error() == "http: no such file" {
			panic(exception.CustomEror{Code: 400, Error: "Foto is required"})
		}
		panic(exception.CustomEror{Code: 400, Error: err.Error()})
	}
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	helper.PanicIfError(err)
	registerRequest.Foto = fileContents

	userRespone := controller.UserService.CreateUser(request.Context(), registerRequest)

	webRespone := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   userRespone,
	}

	helper.WriteToResponseBody(writer, webRespone)

}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	loginRequest := webrequest.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	userLogin := controller.UserService.Login(request.Context(), loginRequest)

	webResponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   userLogin,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (c *UserControllerImpl) Authenticate(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	id, ok := request.Context().Value("id").(int)

	if !ok {
		http.Error(writer, "failed to get valueOne", http.StatusInternalServerError)
		return
	}

	getUser := c.UserService.Authenticate(request.Context(), id)

	webResponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   getUser,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (c *UserControllerImpl) ListAllUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	users := c.UserService.ListAllUsers(request.Context())

	webresponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   users,
	}
	helper.WriteToResponseBody(writer, webresponse)
}

func (c *UserControllerImpl) UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	id, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	updateRequest := webrequest.UpdateUserRequest{}

	for key, values := range request.Form {
		if len(values) > 0 {
			switch key {
			case "name":
				updateRequest.Name = values[0]
			case "email":
				updateRequest.Email = values[0]
			case "gender":
				if values[0] == "male" {
					updateRequest.Gender = "male"
				} else {
					updateRequest.Gender = "female"
				}
			case "telp":
				updateRequest.Telp = values[0]
			case "birthdate":

				if values[0] != "" {

					layout := "2006-01-02"
					parsedTime, err := time.Parse(layout, request.FormValue("birthdate"))
					if err != nil {

						helper.PanicIfError(err)
					}
					updateRequest.Birthdate.Time = parsedTime
					updateRequest.Birthdate.Valid = true
				}
			case "address":
				updateRequest.Address = values[0]
			case "level":
				if values[0] == "member" {
					updateRequest.Level = "member"
				} else if values[0] == "admin" {
					updateRequest.Level = "admin"
				} else {
					updateRequest.Level = "superadmin"
				}
			case "batas":
				updateRequest.Batas = values[0]

			}
		}
	}

	file, _, err := request.FormFile("foto")

	if err != nil {
		fmt.Println("no file")
	}

	var foto []byte

	if file != nil {
		fileContents, err := io.ReadAll(file)
		helper.PanicIfError(err)
		foto = fileContents
		defer file.Close()
	}

	updateRequest.Foto = foto

	_ = c.UserService.UpdateUser(request.Context(), updateRequest, id)

	webresponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   true,
	}
	helper.WriteToResponseBody(writer, webresponse)
}

func (c *UserControllerImpl) FindUserById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	getUser := c.UserService.FindByid(request.Context(), id)
	webresponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   getUser,
	}
	helper.WriteToResponseBody(writer, webresponse)
}
