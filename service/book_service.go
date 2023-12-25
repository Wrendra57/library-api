package service

import (
	"context"

	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
)

type BookService interface {
	CreateBook(ctx context.Context, request webrequest.BookCreateRequest) webresponse.BookResponse
}
