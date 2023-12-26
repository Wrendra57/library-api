package service

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/helper/konversi"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
	"github.com/be/perpustakaan/repository"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository      repository.BookRepository
	CategoryRepository  repository.CategoryBookRepository
	AuthorRepository    repository.AuthorRepository
	PublisherRepository repository.PublisherRepository
	UserRepository      repository.UserRepository
	RakRepository       repository.RakRepository
	DB                  *sql.DB
	Validate            *validator.Validate
	Cld                 *cloudinary.Cloudinary
}

func NewBookService(bookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate, cld *cloudinary.Cloudinary, catRepo repository.CategoryBookRepository, authorRepo repository.AuthorRepository, publiRepo repository.PublisherRepository, userRepo repository.UserRepository, rakRepo repository.RakRepository) BookService {
	return &BookServiceImpl{
		BookRepository:      bookRepository,
		CategoryRepository:  catRepo,
		AuthorRepository:    authorRepo,
		PublisherRepository: publiRepo,
		UserRepository:      userRepo,
		RakRepository:       rakRepo,
		DB:                  DB,
		Validate:            validate,
		Cld:                 cld,
	}
}

func (s *BookServiceImpl) CreateBook(ctx context.Context, request webrequest.BookCreateRequest) webresponse.BookResponse {
	fmt.Println("serviceCreate")
	admin_id, ok := ctx.Value("id").(int)

	if !ok {
		panic(exception.CustomEror{Code: 400, Error: "user not found "})
	}

	err := s.Validate.Struct(request)
	helper.PanicIfError(err)
	book := domain.Book{}

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book.Title = request.Title
	// cek user
	user, err := s.UserRepository.FindById(ctx, tx, admin_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user unauthorized"})
	}
	book.Admin_id = user.User_id

	// handle category
	c := domain.Category{}
	category, err := s.CategoryRepository.FindByName(ctx, tx, request.Category)
	if err != nil {
		c.Category = request.Category
		category = s.CategoryRepository.Create(ctx, tx, c)
	}
	fmt.Println("Category ==>", category)
	book.Category_id = category.Category_id

	// handle author_id
	author, err := s.AuthorRepository.FindByName(ctx, tx, request.Author)
	if err != nil {
		a := domain.Author{
			Name: request.Author,
		}
		author = s.AuthorRepository.Create(ctx, tx, a)
	}
	fmt.Println("author =>>>", author)
	book.Author_id = author.Author_id

	// handle publisher_id
	publisher, err := s.PublisherRepository.FindByName(ctx, tx, request.Publisher)
	if err != nil {
		p := domain.Publisher{
			Name: request.Publisher,
		}
		publisher = s.PublisherRepository.Create(ctx, tx, p)
	}
	fmt.Println("publisher =>>>", publisher)
	book.Publisher_id = publisher.Publisher_id

	book.Isbn = request.Isbn

	page_count, err := strconv.Atoi(request.Page_count)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "page count must number"})
	}
	if page_count <= 0 {
		panic(exception.CustomEror{Code: 400, Error: "page count must greater than zero"})
	}
	book.Page_count = page_count

	stock, err := strconv.Atoi(request.Stock)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "stock must number"})
	}
	if stock <= 0 {
		panic(exception.CustomEror{Code: 400, Error: "stock must greater than zero"})
	}
	book.Stock = stock

	publisher_year, err := strconv.Atoi(request.Publication_year)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "publisher year must number"})
	}
	if publisher_year <= 0 {
		panic(exception.CustomEror{Code: 400, Error: "publisher_year must greater than zero"})
	}
	book.Publication_year = publisher_year

	// handle rak
	rakReq := webrequest.RakByNameRowRequest{
		Name: request.Rak,
	}
	col, err := strconv.Atoi(request.Column)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "column must number"})
	}
	if col <= 0 {
		panic(exception.CustomEror{Code: 400, Error: "column must greater than zero"})
	}
	rakReq.Col = col

	row, err := strconv.Atoi(request.Rows)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "Rows must number"})
	}
	if col <= 0 {
		panic(exception.CustomEror{Code: 400, Error: "Rows must greater than zero"})
	}
	rakReq.Rows_rak = row

	rak, err := s.RakRepository.FindByNameColRow(ctx, tx, rakReq)
	if err != nil {
		a := domain.Rak{
			Name:     request.Rak,
			Rows_rak: row,
			Col:      col,
		}
		rak = s.RakRepository.Create(ctx, tx, a)
	}
	fmt.Println("rakk==>", rak)
	book.Rak_id = rak.Rak_id

	price, err := strconv.Atoi(request.Price)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "price must number"})
	}
	if stock <= 0 {
		panic(exception.CustomEror{Code: 400, Error: "price must greater than zero"})
	}
	book.Price = price

	// handle foto
	reader := bytes.NewReader(request.Foto)

	result, err := s.Cld.Upload.Upload(ctx, reader, uploader.UploadParams{})
	if err != nil {
		fmt.Println(err)
		panic("upload fatal")
	}
	fmt.Println(result.SecureURL)
	book.Foto = result.SecureURL

	fmt.Print("book")
	fmt.Print(book)

	boks := s.BookRepository.Create(ctx, tx, book)
	fmt.Println(boks)
	sads := webresponse.BookResponse{}
	return sads
}

func (s *BookServiceImpl) FindBookById(ctx context.Context, id int) webresponse.BookResponseComplete {
	fmt.Println("service find book")

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getBook, err := s.BookRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "book not found"})
	}
	fmt.Println(getBook)

	return getBook

}

func (s *BookServiceImpl) ListBook(ctx context.Context, request webrequest.FindAllRequest) []webresponse.BookResponseComplete {
	fmt.Println("service list book")
	// deffault limit
	limit := 30
	offset := 0
	maxLimit := 100

	if request.Limit != "" {
		res, err := strconv.Atoi(request.Limit)
		if err == nil {
			if res > maxLimit {
				limit = maxLimit
			} else {
				limit = res
			}
		}
	}
	if request.Offset != "" {
		res, err := strconv.Atoi(request.Offset)
		if err == nil {
			if res >= 0 {
				offset = res
			}
		}
	}

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getBooks := s.BookRepository.ListBook(ctx, tx, limit, offset)

	fmt.Println(getBooks)

	return getBooks
}

func (s *BookServiceImpl) SearchBook(ctx context.Context, search string, l webrequest.FindAllRequest) []webresponse.BookResponseComplete {
	fmt.Println("service list book")
	// deffault limit
	limit := 30
	offset := 0
	maxLimit := 100
	searched := ""

	if l.Limit != "" {
		res, err := strconv.Atoi(l.Limit)
		if err == nil {
			if res > maxLimit {
				limit = maxLimit
			} else {
				limit = res
			}
		}
	}
	if l.Offset != "" {
		res, err := strconv.Atoi(l.Offset)
		if err == nil {
			if res >= 0 {
				offset = res
			}
		}
	}
	if search != "" {
		searched = search
	}
	req := webrequest.SearchBookRequest{
		Search: searched,
		Limit:  limit,
		Offset: offset,
	}

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getBooks := s.BookRepository.FindBook(ctx, tx, req)

	fmt.Println(getBooks)

	return getBooks
}

func (s *BookServiceImpl) UpdateUser(ctx context.Context, request webrequest.UpdateBookRequest, id int) int {
	fmt.Println("service update")
	admin_id, ok := ctx.Value("id").(int)

	if !ok {
		panic(exception.CustomEror{Code: 400, Error: "user not found"})
	}

	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	book := domain.Book{}

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.FindById(ctx, tx, admin_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user unauthorized"})
	}

	// cek book
	_, err = s.BookRepository.FindById(ctx, tx, id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	book.Title = request.Title
	book.Admin_id = user.User_id

	if request.Category != "" {
		c := domain.Category{}
		category, err := s.CategoryRepository.FindByName(ctx, tx, request.Category)
		if err != nil {
			c.Category = request.Category
			category = s.CategoryRepository.Create(ctx, tx, c)
		}
		fmt.Println("Category ==>", category)
		book.Category_id = category.Category_id
	}

	if request.Author != "" {
		author, err := s.AuthorRepository.FindByName(ctx, tx, request.Author)
		if err != nil {
			a := domain.Author{
				Name: request.Author,
			}
			author = s.AuthorRepository.Create(ctx, tx, a)
		}
		fmt.Println("author =>>>", author)
		book.Author_id = author.Author_id
	}

	if request.Publisher != "" {
		publisher, err := s.PublisherRepository.FindByName(ctx, tx, request.Publisher)
		if err != nil {
			p := domain.Publisher{
				Name: request.Publisher,
			}
			publisher = s.PublisherRepository.Create(ctx, tx, p)
		}
		fmt.Println("publisher =>>>", publisher)
		book.Publisher_id = publisher.Publisher_id
	}

	book.Isbn = request.Isbn

	if request.Page_count != "" {
		book.Page_count = konversi.StrToInt(request.Page_count, "page_count")
	}

	if request.Stock != "" {
		book.Stock = konversi.StrToInt(request.Stock, "stock")
	}
	if request.Publication_year != "" {
		book.Publication_year = konversi.StrToInt(request.Publication_year, "publication year")
	}
	if request.Rak != "" && request.Rows != "" && request.Column != "" {
		rakReq := webrequest.RakByNameRowRequest{
			Name: request.Rak,
		}

		rakReq.Col = konversi.StrToInt(request.Column, "column")

		rakReq.Rows_rak = konversi.StrToInt(request.Rows, "rows rak")

		rak, err := s.RakRepository.FindByNameColRow(ctx, tx, rakReq)
		if err != nil {
			a := domain.Rak{
				Name:     request.Rak,
				Rows_rak: rakReq.Col,
				Col:      rakReq.Rows_rak,
			}
			rak = s.RakRepository.Create(ctx, tx, a)
		}
		fmt.Println("rakk==>", rak)
		book.Rak_id = rak.Rak_id
	}

	if request.Price != "" {
		book.Price = konversi.StrToInt(request.Price, "price")
	}
	// handle foto
	if request.Foto != nil {
		reader := bytes.NewReader(request.Foto)

		result, err := s.Cld.Upload.Upload(ctx, reader, uploader.UploadParams{})
		if err != nil {
			fmt.Println(err)
			panic(exception.CustomEror{Code: 400, Error: "gagal upload foto"})
		}
		fmt.Println(result.SecureURL)
		book.Foto = result.SecureURL
	}

	update := s.BookRepository.Update(ctx, tx, id, book)

	fmt.Print("book")
	fmt.Print(book)

	// panic("sda")
	return update
}
