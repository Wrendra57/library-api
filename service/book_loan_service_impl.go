package service

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/helper/compare"
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
	Penalties          repository.PenaltiesRepository
	DB                 *sql.DB
	Validate           *validator.Validate
	Cld                *cloudinary.Cloudinary
}

func NewBookLoanService(bookLoanRepository repository.BookLoanRepository, userRepository repository.UserRepository, bookRepository repository.BookRepository, penaltiesRepository repository.PenaltiesRepository, DB *sql.DB, validate *validator.Validate, cld *cloudinary.Cloudinary) BookLoanService {
	return &BookLoanServiceImpl{
		BookLoanRepository: bookLoanRepository,
		UserRepository:     userRepository,
		BookRepository:     bookRepository,
		Penalties:          penaltiesRepository,
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

	bookLoan.User_id = request.User_id
	bookLoan.Book_id = request.Book_id
	bookLoan.Status = "onloan"
	_, err = s.BookLoanRepository.FindByUserIdBookId(ctx, tx, bookLoan)
	if err == nil {
		panic(exception.CustomEror{Code: 400, Error: "User was rent this book"})
	}

	admin, err := s.UserRepository.FindById(ctx, tx, admin_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "admin unauthorized"})
	}
	bookLoan.Admin_id = admin.User_id

	user, err := s.UserRepository.FindById(ctx, tx, request.User_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user unauthorized"})
	}
	if user.Batas <= 0 {
		panic(exception.CustomEror{Code: 400, Error: "limit for rent"})
	}

	book, err := s.BookRepository.FindById(ctx, tx, request.Book_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "Book Not Found"})
	}
	if book.Stock <= 0 {
		panic(exception.CustomEror{Code: 400, Error: "Stock Book empty"})
	}

	bookLoan.Checkout_date = time.Now()
	bookLoan.Due_date = bookLoan.Checkout_date.Add(3 * 24 * time.Hour)

	updateBook := domain.Book{
		Stock:         book.Stock - 1,
		UpdateForRent: "true",
	}
	_ = s.BookRepository.Update(ctx, tx, book.Book_id, updateBook)

	updateUser := webrequest.UpdateUserRequest{
		Batas: strconv.Itoa(user.Batas - 1),
	}
	_ = s.UserRepository.Update(ctx, tx, user.User_id, updateUser)

	loan := s.BookLoanRepository.Create(ctx, tx, bookLoan)

	resp := helper.ToBookLoanResponse(loan)
	return resp
}

func (s *BookLoanServiceImpl) ReturnBookLoan(ctx context.Context, request webrequest.BookLoanCreateRequest) webresponse.BookLoanResponseComplete2 {
	admin_id, ok := ctx.Value("id").(int)
	timeNow := time.Now()
	if !ok {
		panic(exception.CustomEror{Code: 400, Error: "user not found "})
	}
	fmt.Println(request.Id)

	err := s.Validate.Struct(request)
	helper.PanicIfError(err)
	bookLoan := domain.BookLoan{}

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	bookLoan.User_id = request.User_id
	bookLoan.Book_id = request.Book_id
	bookLoan.Status = "onloan"

	if request.Id != 0 {
		bookLoan, err = s.BookLoanRepository.FindById(ctx, tx, request.Id)
		if err != nil {
			panic(exception.CustomEror{Code: 400, Error: "User not rent this book"})
		}
	} else {
		bookLoan, err = s.BookLoanRepository.FindByUserIdBookId(ctx, tx, bookLoan)
		if err != nil {
			panic(exception.CustomEror{Code: 400, Error: "User not rent this book"})
		}
	}

	bookLoan.Return_date.Time = time.Now()
	bookLoan.Return_date.Valid = true

	admin, err := s.UserRepository.FindById(ctx, tx, admin_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "admin not found"})
	}
	user, err := s.UserRepository.FindById(ctx, tx, bookLoan.User_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user not found"})
	}
	book, err := s.BookRepository.FindById(ctx, tx, bookLoan.Book_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "Book Not Found"})
	}

	updateBookLoan := webrequest.BookLoanUpdateRequest{
		Loan_id:  bookLoan.Loan_id,
		Admin_id: admin_id,
	}
	updateBookLoan.Return_date.Time = timeNow
	updateBookLoan.Return_date.Valid = true

	penalty := domain.Penalties{}
	if compare.CompareTime(timeNow.Format(time.RFC3339), bookLoan.Due_date.Format(time.RFC3339)) {
		updateBookLoan.Status = "returned"
		bookLoan.Status = "returned"
		fmt.Println("not late")
	} else {
		updateBookLoan.Status = "overdue"
		bookLoan.Status = "overdue"

		penalty.Admin_id = admin_id
		penalty.Due_date = time.Now().Add(3 * 24 * time.Hour)
		penalty.Loan_id = bookLoan.Loan_id
		penalty.Payment_status = "unpaid"
		duration := math.Ceil(timeNow.Sub(bookLoan.Due_date).Hours() / 24)
		penalty.Penalty_amount = 5_000 * int(duration)
		penalty.Reason = "late"

		penalty = s.Penalties.Create(ctx, tx, penalty)
	}

	_ = s.BookLoanRepository.Update(ctx, tx, updateBookLoan)

	_ = s.UserRepository.Update(ctx, tx, user.User_id, webrequest.UpdateUserRequest{
		Batas: strconv.Itoa(user.Batas + 1),
	})

	_ = s.BookRepository.Update(ctx, tx, book.Book_id, domain.Book{
		Stock:         book.Stock + 1,
		UpdateForRent: "true",
	})

	res := helper.ToBookLoanResponseComplete(bookLoan, book, user, admin, penalty)
	return res
}

func (s *BookLoanServiceImpl) FindAll(ctx context.Context, request webrequest.ListALlBookLoanRequest) []webresponse.ListBookLoanResponse {
	limit := 10
	offset := 0

	if request.Limit > 0 {
		limit = request.Limit
	}
	if request.Offset > 0 {
		offset = request.Offset - 1
	}

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getList := s.BookLoanRepository.FindAll(ctx, tx, limit, offset)

	return getList
}

func (s *BookLoanServiceImpl) FindById(ctx context.Context, id int) webresponse.ListBookLoanResponse {
	user_id, ok := ctx.Value("id").(int)

	if !ok {
		panic(exception.CustomEror{Code: 400, Error: "user not found "})
	}

	level, ok := ctx.Value("level").(string)
	if !ok {
		panic(exception.CustomEror{Code: 400, Error: "user not found "})
	}

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	bookLoan, err := s.BookLoanRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	if level == "member" {
		if user_id != bookLoan.User_id {
			panic(exception.CustomEror{Code: 400, Error: "Not restricted access"})
		}
	}
	book, err := s.BookRepository.FindById(ctx, tx, bookLoan.Book_id)
	helper.PanicIfError(err)

	user, err := s.UserRepository.FindById(ctx, tx, bookLoan.User_id)
	helper.PanicIfError(err)

	penalty, err := s.Penalties.FindById(ctx, tx, bookLoan.Loan_id)

	res := helper.ToDetailBookLoanResponseComplete(bookLoan, book, user, penalty)
	return res
	// panic("s")
}

func (s *BookLoanServiceImpl) ListByUserId(ctx context.Context) []webresponse.ListBookLoanResponse {
	user_id, ok := ctx.Value("id").(int)
	if !ok {
		panic(exception.CustomEror{Code: 400, Error: "user not found "})
	}

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	list, err := s.BookLoanRepository.ListByUserId(ctx, tx, user_id)
	helper.PanicIfError(err)

	return list

}
