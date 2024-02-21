package domain

import (
	"database/sql"
	"time"
)

type BookLoan struct {
	Loan_id       int
	Checkout_date time.Time
	Due_date      time.Time
	Return_date   sql.NullTime
	Status        string
	Book_id       int
	User_id       int
	Admin_id      int
	Created_at    time.Time
	Updated_at    time.Time
}
