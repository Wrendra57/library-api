package service

import (
	"context"

	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
)

type BookService interface {
	CreateBook(ctx context.Context, request webrequest.BookCreateRequest) webresponse.BookResponseComplete

	FindBookById(ctx context.Context, id int) webresponse.BookResponseComplete
	ListBook(ctx context.Context, request webrequest.FindAllRequest) []webresponse.BookResponseComplete
	SearchBook(ctx context.Context, search string, limit webrequest.FindAllRequest) []webresponse.BookResponseComplete
	UpdateBook(ctx context.Context, request webrequest.UpdateBookRequest, id int) int
	DeleteBook(ctx context.Context, id int) int
}
