package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
)

type AuthorRepository interface {
	Create(ctx context.Context, tx *sql.Tx, author domain.Author) domain.Author
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Author, error)
	FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Author, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Author
	// Update(ctx context.Context, tx *sql.Tx, id int )
}
