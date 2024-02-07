package webresponse

import (
	"database/sql"
	"time"

	"github.com/be/perpustakaan/model/domain"
)

type BookLoanResponseComplete struct {
	Loan_id       int          `json:"loan_id"`
	Checkout_date time.Time    `json:"checkout_date"`
	Due_date      time.Time    `json:"due_date"`
	Return_date   sql.NullTime `json:"return_date"`
	Status        string       `json:"status"`
	Book_id       int          `json:"book_id"`
	User_id       int          `json:"user_id"`
	Admin_id      int          `json:"admin_id"`
	Created_at    time.Time    `json:"created_at"`
	Updated_at    time.Time    `json:"updated_at"`
}

type BookLoanResponseComplete2 struct {
	Loan_id       int                  `json:"loan_id"`
	Checkout_date time.Time            `json:"checkout_date"`
	Due_date      time.Time            `json:"due_date"`
	Return_date   sql.NullTime         `json:"return_date"`
	Status        string               `json:"status"`
	Book_id       BookResponseComplete `json:"book_id"`
	User_id       domain.User          `json:"user"`
	Admin_id      domain.User          `json:"admin_id"`
	Penalties     domain.Penalties     `json:"penalties"`
	Created_at    time.Time            `json:"created_at"`
	Updated_at    time.Time            `json:"updated_at"`
}
