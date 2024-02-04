package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
)

type BookLoanRepositoryImpl struct {
}

func NewBookLoanRepository() BookLoanRepository {
	return &BookLoanRepositoryImpl{}
}

func (r *BookLoanRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, l domain.BookLoan) domain.BookLoan {
	SQL := `INSERT INTO book_loan(checkout_date,due_date,status,book_id,user_id,admin_id) values(?,?,?,?,?,?)`

	result, err := tx.ExecContext(ctx, SQL, l.Checkout_date, l.Due_date, l.Status, l.Book_id, l.User_id, l.Admin_id)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	l.Loan_id = int(id)
	l.Created_at = time.Now()
	l.Updated_at = time.Now()

	return l
}

func (r *BookLoanRepositoryImpl) FindByUserIdBookId(ctx context.Context, tx *sql.Tx, userId int, bookId int) (domain.BookLoan, error) {
	SQL := `select loan_id, checkout_date, due_date, return_date, status, book_id, user_id, admin_id, created_at, updated_at from book_loan where user_id=? and book_id=? and status='onloan'`

	rows, err := tx.QueryContext(ctx, SQL, userId, bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	l := domain.BookLoan{}
	if rows.Next() {
		err := rows.Scan(&l.Loan_id, &l.Checkout_date, &l.Due_date, &l.Return_date, &l.Status, &l.Book_id, &l.User_id, &l.Admin_id, &l.Created_at, &l.Updated_at)
		helper.PanicIfError(err)
		return l, nil
	} else {
		return l, errors.New("user not found")
	}
}
