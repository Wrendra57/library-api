package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
)

type BookRepository interface {
	Create(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
}
