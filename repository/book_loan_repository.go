package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
)

type BookLoanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, loan domain.BookLoan) domain.BookLoan
	FindByUserIdBookId(ctx context.Context, tx *sql.Tx, req domain.BookLoan) (domain.BookLoan, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.BookLoan, error)
	Update(ctx context.Context, tx *sql.Tx, req webrequest.BookLoanUpdateRequest) int
	FindAll(ctx context.Context, tx *sql.Tx, limit int, offset int) []webresponse.ListBookLoanResponse
}
