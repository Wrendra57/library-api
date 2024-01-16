package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
)

type BookLoanRepositoryImpl struct {
}

func NewBookLoanRepository() BookLoanRepository {
	return &BookLoanRepositoryImpl{}
}

func (r *BookLoanRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, loan domain.Loan) domain.Loan {
	// SQL:= `
	// 	insert into book_loan(checkout_date, due_date, status, book_id, user_id, admin_id)
	// `

	panic("das")
}
