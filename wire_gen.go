// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/be/perpustakaan/app"
	"github.com/be/perpustakaan/controller"
	"github.com/be/perpustakaan/repository"
	"github.com/be/perpustakaan/service"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	userRepository := repository.NewUserRepository()
	db := app.NewDB()
	validate := app.NewValidate()
	cloudinary := app.Cloudinary()
	userService := service.NewUserService(userRepository, db, validate, cloudinary)
	userController := controller.NewUserController(userService)
	bookRepository := repository.NewBookRepository()
	categoryBookRepository := repository.NewCategoryRepository()
	authorRepository := repository.NewAuthorRepository()
	publisherRepository := repository.NewPublisherRepository()
	rakRepository := repository.NewRakRepository()
	bookService := service.NewBookService(bookRepository, db, validate, cloudinary, categoryBookRepository, authorRepository, publisherRepository, userRepository, rakRepository)
	bookController := controller.NewBookController(bookService)
	bookLoanRepository := repository.NewBookLoanRepository()
	penaltiesRepository := repository.NewPenaltiesRepository()
	bookLoanService := service.NewBookLoanService(bookLoanRepository, userRepository, bookRepository, penaltiesRepository, db, validate, cloudinary)
	bookLoanController := controller.NewBookLoanController(bookLoanService)
	router := app.NewRouter(userController, bookController, bookLoanController)
	server := NewServer(router)
	return server
}
