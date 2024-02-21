package exception

import (
	"fmt"
	"net/http"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/webresponse"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationErrors(writer, request, err) {
		return
	}
	if duplicateEmailEror(writer, request, err) {
		return
	}
	if customEror(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func customErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s must be maximum %s characters long", err.Field(), err.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	case "oneof":
		return fmt.Sprintf("%s must be a %s", err.Field(), err.Param())
	default:
		return fmt.Sprintf("Validation error on field %s with tag %s", err.Field(), err.Tag())
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	// fmt.Println(exception)
	if ok {

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		for _, e := range err.(validator.ValidationErrors) {
			// fmt.Println("ajlaskd")
			// fmt.Println("Field:", e.Field())
			// fmt.Println("Error:", customErrorMessage(e))

			webResponse := webresponse.ResponseApi{
				Code:   http.StatusBadRequest,
				Status: customErrorMessage(e),
				Data:   nil,
			}
			helper.WriteToResponseBody(writer, webResponse)
			return true
		}

		return true
	} else {
		return false
	}

}

func duplicateEmailEror(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(DuplicateEmailError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := webresponse.ResponseApi{
			Code:   http.StatusBadRequest,
			Status: exception.Error,
			Data:   nil,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func customEror(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(CustomEror)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(exception.Code)

		webResponse := webresponse.ResponseApi{
			Code:   exception.Code,
			Status: exception.Error,
			Data:   nil,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	// fmt.Println("err")
	webResponse := webresponse.ResponseApi{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
