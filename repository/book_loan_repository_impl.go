package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
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

func (r *BookLoanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, limit int, offset int) []webresponse.ListBookLoanResponse {
	SQL := `select 
				bl.loan_id, 
				bl.checkout_date, 
				bl.due_date, 
				bl.return_date, 
				bl.status, 
				bl.admin_id,
				JSON_OBJECT(
					"book_id", b.book_id,
					"book_title", b.title,
					"foto", b.foto
				) as book,
				JSON_OBJECT(
					"user_id", u.user_id,
					"name", u.name,
					"foto", u.foto
				) as users,
				bl.created_at,
				JSON_OBJECT(
					"penalty_id", p.penalty_id,
					"penalty_amount", p.penalty_amount,
					"payment_status", p.payment_status,
					"reason", p.reason
				) as penalty
				FROM library.book_loan as bl
			LEFT JOIN library.penalties as p on p.loan_id= bl.loan_id
			LEFT JOIN library.book as b on b.book_id=bl.book_id
			LEFT JOIN library.user as u on u.user_id=bl.user_id
			ORDER BY bl.updated_at DESC
			LIMIT ? 
			OFFSET ?`

	rows, err := tx.QueryContext(ctx, SQL, limit, offset)
	helper.PanicIfError(err)
	defer rows.Close()

	var bookLoan []webresponse.ListBookLoanResponse
	for rows.Next() {
		b := webresponse.ListBookLoanResponse{}
		var bookJSON, userJSON, penaltyJSON []byte

		err := rows.Scan(&b.Loan_id, &b.Checkout_date, &b.Due_date, &b.Return_date, &b.Status, &b.Admin_id, &bookJSON, &userJSON, &b.Created_at, &penaltyJSON)
		helper.PanicIfError(err)

		var book webresponse.Book
		err = json.Unmarshal(bookJSON, &book)
		helper.PanicIfError(err)

		var user webresponse.User
		err = json.Unmarshal(userJSON, &user)
		helper.PanicIfError(err)
		fmt.Println("penaltyJSON")
		fmt.Println(penaltyJSON)

		var penalty webresponse.Penalty
		err = json.Unmarshal(penaltyJSON, &penalty)
		helper.PanicIfError(err)
		fmt.Println(penalty)

		b.Book = book
		b.User = user
		b.Penalties = penalty
		bookLoan = append(bookLoan, b)
	}
	return bookLoan
}
