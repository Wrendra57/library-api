package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
)

type RakRepository interface {
	Create(ctx context.Context, tx *sql.Tx, rak domain.Rak) domain.Rak
	FindByNameColRow(ctx context.Context, tx *sql.Tx, rak webrequest.RakByNameRowRequest) (domain.Rak, error)
}
