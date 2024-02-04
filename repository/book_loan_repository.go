package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
)

type BookLoanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, loan domain.BookLoan) domain.BookLoan
	FindByUserIdBookId(ctx context.Context, tx *sql.Tx, userId int, bookId int) (domain.BookLoan, error)
}
