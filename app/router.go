package app

import (
	"github.com/be/perpustakaan/controller"
	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/middleware"

	// _ "github.com/be/perpustakaan/middleware"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController, bookController controller.BookController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.GET("/api/user", middleware.AuthMiddleware(userController.Authenticate))
	router.GET("/api/users", middleware.AuthMiddleware(middleware.RoleMiddleware("admin", userController.ListAllUsers)))
	router.PUT("/api/user/:id", middleware.AuthMiddleware(userController.UpdateUser))
	router.GET("/api/user/:id", middleware.AuthMiddleware(middleware.RoleMiddleware("admin", userController.FindUserById)))

	// Books
	router.POST("/api/book", middleware.AuthMiddleware(middleware.RoleMiddleware("admin", bookController.CreateBook)))
	router.GET("/api/book/:id", bookController.FindBookById)
	router.GET("/api/books", bookController.ListBooks)
	router.GET("/api/books/search", bookController.SearchBook)
	router.PUT("/api/books/:id", middleware.AuthMiddleware(middleware.RoleMiddleware("admin", bookController.UpdateBook)))
	router.DELETE("/api/books/:id", middleware.AuthMiddleware(middleware.RoleMiddleware("admin", bookController.DeleteBook)))
	// router.GET("/api/user", middleware.AuthMiddleware(userController.Authenticate))
	// router.GET("/api/users", middleware.AuthMiddleware(middleware.RoleMiddleware("member", userController.Authenticate)))

	router.PanicHandler = exception.ErrorHandler
	return router
}
