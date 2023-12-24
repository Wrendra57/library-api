package controller

import (
	"fmt"
	"io"
	"net/http"
	"time"

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

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	err := request.ParseMultipartForm(10<<20)
	helper.PanicIfError(err)
	registerRequest := webrequest.UserCreateRequest{}

	registerRequest.Name = request.FormValue("name")
	registerRequest.Email = request.FormValue("email")
	registerRequest.Password = request.FormValue("password")
	registerRequest.Gender = request.FormValue("gender")
	registerRequest.Telp = request.FormValue("telp")

	// parsing date
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, request.FormValue("birthdate"))
	if err != nil {
		fmt.Println("Error:", err)
		helper.PanicIfError(err)
	}
	fmt.Println(parsedTime)
	registerRequest.Birthdate = parsedTime
	registerRequest.Address = request.FormValue("address")

	file, _, err := request.FormFile("foto")
	helper.PanicIfError(err)
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	helper.PanicIfError(err)
	registerRequest.Foto = fileContents	
	
	userRespone := controller.UserService.CreateUser(request.Context(),registerRequest)

	webRespone := webresponse.ResponseApi{
		Code: 200,
		Status: "OK",
		Data: userRespone,
	}

	helper.WriteToResponseBody(writer,webRespone)


}

