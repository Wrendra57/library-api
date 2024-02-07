package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
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

func (r *BookLoanRepositoryImpl) FindByUserIdBookId(ctx context.Context, tx *sql.Tx, req domain.BookLoan) (domain.BookLoan, error) {
	SQL := `select loan_id, checkout_date, due_date, return_date, status, book_id, user_id, admin_id, created_at, updated_at from book_loan where user_id=? and book_id=?`

	var args []interface{}
	args = append(args, req.User_id)
	args = append(args, req.Book_id)

	if req.Status != "" {
		SQL += " and status = ? "
		args = append(args, req.Status)
	}
	rows, err := tx.QueryContext(ctx, SQL, args...)
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

func (r *BookLoanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.BookLoan, error) {
	SQL := `select loan_id, checkout_date, due_date, return_date, status, book_id, user_id, admin_id, created_at, updated_at from book_loan where loan_id= ?`
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	fmt.Println(id)
	defer rows.Close()

	l := domain.BookLoan{}
	if rows.Next() {
		err := rows.Scan(&l.Loan_id, &l.Checkout_date, &l.Due_date, &l.Return_date, &l.Status, &l.Book_id, &l.User_id, &l.Admin_id, &l.Created_at, &l.Updated_at)
		helper.PanicIfError(err)
		return l, nil
	} else {
		return l, errors.New("book loan not found")
	}
}

func (r *BookLoanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, req webrequest.BookLoanUpdateRequest) int {
	SQL := "update book_loan set  "
	var args []interface{}

	if req.Return_date.Valid {
		SQL += "return_date = ?, "
		args = append(args, req.Return_date.Time)
	}
	if req.Status != "" {
		SQL += "status = ?, "
		args = append(args, req.Status)
	}
	if req.Admin_id != 0 {
		SQL += "admin_id = ?  "
		args = append(args, req.Admin_id)
	}
	SQL = SQL[:len(SQL)-2]

	SQL += " WHERE loan_id = ?"
	args = append(args, req.Loan_id)
	fmt.Println(SQL)
	_, err := tx.Exec(SQL, args...)
	helper.PanicIfError(err)

	return req.Loan_id
}
