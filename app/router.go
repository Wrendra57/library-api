package app

import (
	"github.com/be/perpustakaan/controller"
	"github.com/be/perpustakaan/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", userController.Register)
	// router.POST("/api/users/login", userController.Login)

	router.PanicHandler = exception.ErrorHandler
	return router
}