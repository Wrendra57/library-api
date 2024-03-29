package service

import (
	"bytes"
	"context"
	"database/sql"
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

func (s *BookServiceImpl) CreateBook(ctx context.Context, request webrequest.BookCreateRequest) webresponse.BookResponseComplete {

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

	// cek user
	user, err := s.UserRepository.FindById(ctx, tx, admin_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user unauthorized"})
	}
	book.Admin_id = user.User_id
	book.Title = request.Title

	// handle category
	c := domain.Category{}
	category, err := s.CategoryRepository.FindByName(ctx, tx, request.Category)
	if err != nil {
		c.Category = request.Category
		category = s.CategoryRepository.Create(ctx, tx, c)
	}

	book.Category_id = category.Category_id

	// handle author_id
	author, err := s.AuthorRepository.FindByName(ctx, tx, request.Author)
	if err != nil {
		a := domain.Author{
			Name: request.Author,
		}
		author = s.AuthorRepository.Create(ctx, tx, a)
	}

	book.Author_id = author.Author_id

	// handle publisher_id
	publisher, err := s.PublisherRepository.FindByName(ctx, tx, request.Publisher)
	if err != nil {
		p := domain.Publisher{
			Name: request.Publisher,
		}
		publisher = s.PublisherRepository.Create(ctx, tx, p)
	}

	book.Publisher_id = publisher.Publisher_id

	book.Isbn = request.Isbn

	page_count := konversi.StrToInt(request.Page_count, "page_count")
	book.Page_count = page_count

	stock := konversi.StrToInt(request.Stock, "stock")
	book.Stock = stock

	publisher_year := konversi.StrToInt(request.Publication_year, "publication_year")
	book.Publication_year = publisher_year

	// handle rak
	rakReq := webrequest.RakByNameRowRequest{
		Name: request.Rak,
	}
	col := konversi.StrToInt(request.Column, "column")
	rakReq.Col = col

	row := konversi.StrToInt(request.Rows, "rows")
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

	book.Rak_id = rak.Rak_id

	price := konversi.StrToInt(request.Price, "price")
	book.Price = price

	// handle foto
	reader := bytes.NewReader(request.Foto)

	result, err := s.Cld.Upload.Upload(ctx, reader, uploader.UploadParams{})
	if err != nil {

		panic("upload fatal")
	}

	book.Foto = result.SecureURL

	boks := s.BookRepository.Create(ctx, tx, book)

	resp := webresponse.BookResponseComplete{
		Book_id:          boks.Book_id,
		Title:            boks.Title,
		Category:         category.Category,
		Author:           author.Name,
		Publisher:        publisher.Name,
		Isbn:             boks.Isbn,
		Page_count:       boks.Page_count,
		Stock:            boks.Stock,
		Publication_year: boks.Publication_year,
		Foto:             boks.Foto,
		Rak:              rakReq.Name,
		Column:           rakReq.Col,
		Rows_rak:         rakReq.Rows_rak,
		Price:            boks.Price,
		Admin:            user.Name,
		Created_at:       boks.Created_at,
		Updated_at:       boks.Updated_at,
	}
	return resp
}

func (s *BookServiceImpl) FindBookById(ctx context.Context, id int) webresponse.BookResponseComplete {

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getBook, err := s.BookRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "book not found"})
	}

	return getBook

}

func (s *BookServiceImpl) ListBook(ctx context.Context, request webrequest.FindAllRequest) []webresponse.BookResponseComplete {

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

	return getBooks
}

func (s *BookServiceImpl) SearchBook(ctx context.Context, search string, l webrequest.FindAllRequest) []webresponse.BookResponseComplete {

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

	return getBooks
}

func (s *BookServiceImpl) UpdateBook(ctx context.Context, request webrequest.UpdateBookRequest, id int) int {

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

			panic(exception.CustomEror{Code: 400, Error: "gagal upload foto"})
		}

		book.Foto = result.SecureURL
	}

	update := s.BookRepository.Update(ctx, tx, id, book)

	// panic("sda")
	return update
}

func (s *BookServiceImpl) DeleteBook(ctx context.Context, id int) int {

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = s.BookRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "book not found"})
	}

	err = s.BookRepository.DeleteBook(ctx, tx, id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "failed delete book"})
	}
	return id

}
