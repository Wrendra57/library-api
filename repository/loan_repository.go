package repository

import (
	"context"
	"database/sql"
)

type LoanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, )
}