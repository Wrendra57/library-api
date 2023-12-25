package app

import (
	"github.com/be/perpustakaan/controller"
	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/middleware"

	// _ "github.com/be/perpustakaan/middleware"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.GET("/api/user", middleware.AuthMiddleware(userController.Authenticate))
	router.GET("/api/users", middleware.AuthMiddleware(middleware.RoleMiddleware("admin", userController.ListAllUsers)))
	router.PUT("/api/user/:id",  middleware.AuthMiddleware(userController.UpdateUser))



	// router.GET("/api/user", middleware.AuthMiddleware(userController.Authenticate))
	// router.GET("/api/users", middleware.AuthMiddleware(middleware.RoleMiddleware("member", userController.Authenticate)))

	router.PanicHandler = exception.ErrorHandler
	return router
}