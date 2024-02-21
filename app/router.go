package app

import (
	"net/http"

	"github.com/be/perpustakaan/controller"
	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/middleware"
	"github.com/be/perpustakaan/model/webresponse"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController, bookController controller.BookController, bookLoanController controller.BookLoanController, penaltiesController controller.PenaltiesController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		server := webresponse.ResponseApi{Code: 200, Status: "server already running"}
		helper.WriteToResponseBody(w, server)

	})
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

	// Loan
	router.POST("/api/loan", middleware.AuthMiddleware(middleware.RoleMiddleware("admin", bookLoanController.CreateBookLoan)))
	router.POST("/api/loan/return", middleware.AuthMiddleware(middleware.RoleMiddleware("admin", bookLoanController.ReturnBookLoan)))
	router.GET("/api/loan", middleware.AuthMiddleware(middleware.RoleMiddleware("admin", bookLoanController.FindAll)))
	router.GET("/api/loan/:id", middleware.AuthMiddleware(bookLoanController.FindById))
	router.GET("/api/loans/mylist", middleware.AuthMiddleware(bookLoanController.ListByUserId))

	// Penalties
	router.POST("/api/penalty/pay/:id", middleware.AuthMiddleware(middleware.RoleMiddleware("admin", penaltiesController.PayPenalties)))

	router.PanicHandler = exception.ErrorHandler
	return router
}
