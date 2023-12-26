package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
)

type BookRepository interface {
	Create(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	FindById(ctx context.Context, tx *sql.Tx, id int) (webresponse.BookResponseComplete, error)
	ListBook(ctx context.Context, tx *sql.Tx, limit int, offset int) []webresponse.BookResponseComplete
	FindBook(ctx context.Context, tx *sql.Tx, search webrequest.SearchBookRequest) []webresponse.BookResponseComplete
	Update(ctx context.Context, tx *sql.Tx, id int, book domain.Book) int
}
