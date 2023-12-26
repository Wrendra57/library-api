package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webresponse"
)

type BookRepository interface {
	Create(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	FindById(ctx context.Context, tx *sql.Tx, id int) (webresponse.BookResponseComplete, error)
}
