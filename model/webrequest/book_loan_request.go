package webrequest

import (
	"database/sql"
	"time"
)

type BookLoanCreateRequest struct {
	Id      int `validate:"omitempty,required,number" json:"loan_id"`
	Book_id int `validate:"omitempty,required,number" json:"book_id"`
	User_id int `validate:"omitempty,required,number" json:"user_id"`
}

type BookLoanUpdateRequest struct {
	Loan_id       int          `validate:"omitempty,required,number" json:"loan_id"`
	Checkout_date time.Time    `validate:"omitempty,required,number" json:"checkout_date"`
	Due_date      time.Time    `validate:"omitempty,required,number" json:"due_date"`
	Return_date   sql.NullTime `validate:"omitempty,required,number" json:"return_date"`
	Status        string       `validate:"omitempty,required,number" json:"status"`
	Book_id       int          `validate:"omitempty,required,number" json:"book_id"`
	User_id       int          `validate:"omitempty,required,number" json:"user_id"`
	Admin_id      int          `validate:"omitempty,required,number" json:"admin_id"`
}

type ListALlBookLoanRequest struct {
	Limit  int `validate:"omitempty,required,number" json:"limit"`
	Offset int `validate:"omitempty,required,number" json:"offset"`
}
