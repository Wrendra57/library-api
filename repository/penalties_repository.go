package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
)

type PenaltiesRepository interface {
	Create(ctx context.Context, tx *sql.Tx, penalty domain.Penalties) domain.Penalties
}
