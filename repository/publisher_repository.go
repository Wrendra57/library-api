package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
)

type PublisherRepository interface {
	Create(ctx context.Context, tx *sql.Tx, publisher domain.Publisher) domain.Publisher
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Publisher, error)
	FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Publisher, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Publisher
}
