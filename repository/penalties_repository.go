package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
)

type PenaltiesRepository interface {
	Create(ctx context.Context, tx *sql.Tx, penalty domain.Penalties) domain.Penalties
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Penalties, error)
	Update(ctx context.Context, tx *sql.Tx, id int, penalty webrequest.UpdatePenaltiesRequest) webrequest.UpdatePenaltiesRequest
	// FindAll(ctx context.Context, tx *sql.Tx)
}
