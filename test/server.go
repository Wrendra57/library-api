package test

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/be/perpustakaan/app"
	"github.com/be/perpustakaan/controller"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/repository"
	"github.com/be/perpustakaan/service"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func SetupTestDB() *sql.DB {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/library?parseTime=true")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
func SetupCloudinary() *cloudinary.Cloudinary {
	cld, errCld := cloudinary.NewFromURL("cloudinary://288711213593925:rCvBQ9jbETpNtG_B7Aec13Zel3U@dhtypvjsk")
	helper.PanicIfError(errCld)
	return cld
}
func SetupRouter(db *sql.DB) http.Handler {
	userRepository := repository.NewUserRepository()
	validate := validator.New()
	cld := SetupCloudinary()

	userService := service.NewUserService(userRepository, db, validate, cld)
	userController := controller.NewUserController(userService)
	bookRepository := repository.NewBookRepository()
	categoryBookRepository := repository.NewCategoryRepository()
	authorRepository := repository.NewAuthorRepository()
	publisherRepository := repository.NewPublisherRepository()
	rakRepository := repository.NewRakRepository()
	penaltiesRepository := repository.NewPenaltiesRepository()
	bookService := service.NewBookService(bookRepository, db, validate, cld, categoryBookRepository, authorRepository, publisherRepository, userRepository, rakRepository)
	bookController := controller.NewBookController(bookService)
	bookLoanRepository := repository.NewBookLoanRepository()
	bookLoanService := service.NewBookLoanService(bookLoanRepository, userRepository, bookRepository, penaltiesRepository, db, validate, cld)
	bookLoanController := controller.NewBookLoanController(bookLoanService)
	penaltiesService := service.NewPenaltiesService(userRepository, penaltiesRepository, db, validate)
	penaltisContoller := controller.NewPenaltiesController(penaltiesService)
	router := app.NewRouter(userController, bookController, bookLoanController, penaltisContoller)
	return router
}

func DeleteUser(db *sql.DB) {
	db.Exec("delete from user where email like '%testing%'")
}
