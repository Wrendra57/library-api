package webresponse

import (
	"database/sql"
	"time"
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
