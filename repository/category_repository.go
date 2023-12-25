package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
)

type CategoryBookRepository interface {
	Create(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error)
	FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
