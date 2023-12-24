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

type ContextKey string

const (
	// KeyOne and KeyTwo are keys to access values in the context.
	Id ContextKey = "id"
	Email ContextKey = "email"
	Level ContextKey = "level"
	Token ContextKey = "token"
)
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

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	loginRequest := webrequest.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)
	fmt.Println(loginRequest)
	userLogin := controller.UserService.Login(request.Context(),loginRequest)
	

	// userResponse := controller.UserService.CreateUser(request.Context(), loginRequest )
	webResponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   userLogin,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (c *UserControllerImpl) Authenticate(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
 
	id, ok := request.Context().Value("id").(int)
 
	if !ok {
		http.Error(writer, "failed to get valueOne", http.StatusInternalServerError)
		return
	}
	fmt.Println("ff")
	 
	getUser := c.UserService.Authenticate(request.Context(), id)

	webResponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   getUser,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

