package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
	"github.com/be/perpustakaan/repository"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
)

type BookLoanServiceImpl struct {
	BookLoanRepository repository.BookLoanRepository
	UserRepository     repository.UserRepository
	BookRepository     repository.BookRepository
	DB                 *sql.DB
	Validate           *validator.Validate
	Cld                *cloudinary.Cloudinary
}

func NewBookLoanService(bookLoanRepository repository.BookLoanRepository, userRepository repository.UserRepository, bookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate, cld *cloudinary.Cloudinary) BookLoanService {
	return &BookLoanServiceImpl{
		BookLoanRepository: bookLoanRepository,
		UserRepository:     userRepository,
		BookRepository:     bookRepository,
		DB:                 DB,
		Validate:           validate,
		Cld:                cld,
	}
}

func (s *BookLoanServiceImpl) CreateBookLoan(ctx context.Context, request webrequest.BookLoanCreateRequest) webresponse.BookLoanResponseComplete {
	admin_id, ok := ctx.Value("id").(int)

	if !ok {
		panic(exception.CustomEror{Code: 400, Error: "user not found "})
	}

	err := s.Validate.Struct(request)
	helper.PanicIfError(err)
	bookLoan := domain.BookLoan{}

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = s.BookLoanRepository.FindByUserIdBookId(ctx, tx, request.User_id, request.Book_id)
	if err == nil {
		panic(exception.CustomEror{Code: 400, Error: "User was rent this book"})
	}

	admin, err := s.UserRepository.FindById(ctx, tx, admin_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "admin unauthorized"})
	}
	bookLoan.Admin_id = admin.User_id
	fmt.Println("14")
	user, err := s.UserRepository.FindById(ctx, tx, request.User_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user unauthorized"})
	}
	if user.Batas <= 0 {
		panic(exception.CustomEror{Code: 400, Error: "limit for rent"})
	}
	bookLoan.User_id = user.User_id
	fmt.Println("15")
	book, err := s.BookRepository.FindById(ctx, tx, request.Book_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "Book Not Found"})
	}
	if book.Stock <= 0 {
		panic(exception.CustomEror{Code: 400, Error: "Stock Book empty"})
	}
	bookLoan.Book_id = book.Book_id
	bookLoan.Checkout_date = time.Now()
	bookLoan.Due_date = bookLoan.Checkout_date.Add(3 * 24 * time.Hour)
	bookLoan.Status = "onloan"
	fmt.Println("16")

	updateBook := domain.Book{
		Stock:         book.Stock - 1,
		UpdateForRent: "true",
	}
	_ = s.BookRepository.Update(ctx, tx, book.Book_id, updateBook)
	fmt.Println("1")
	updateUser := webrequest.UpdateUserRequest{
		Batas: strconv.Itoa(user.Batas - 1),
	}
	_ = s.UserRepository.Update(ctx, tx, user.User_id, updateUser)
	fmt.Println("12")
	loan := s.BookLoanRepository.Create(ctx, tx, bookLoan)

	resp := helper.ToBookLoanResponse(loan)
	return resp
}
