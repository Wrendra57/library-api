package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
)

type BookLoanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, loan domain.Loan) domain.Loan
}
