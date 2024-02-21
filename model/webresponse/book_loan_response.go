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
type User struct {
	Name    string `json:"name"`
	Foto    string `json:"foto"`
	User_id int    `json:"user_id"`
}

type Book struct {
	Book_id int    `json:"book_id"`
	Title   string `json:"title"`
	Foto    string `json:"foto"`
}
type Penalty struct {
	Penalty_id     int    `json:"penalty_id"`
	Penalty_amount int    `json:"penalty_amount"`
	Payment_status string `json:"payment_status"`
	Due_date       string `json:"due_date"`
	Reason         string `json:"reason"`
}
type ListBookLoanResponse struct {
	Loan_id       int          `json:"loan_id"`
	Checkout_date time.Time    `json:"checkout_date"`
	Due_date      time.Time    `json:"due_date"`
	Return_date   sql.NullTime `json:"return_date"`
	Status        string       `json:"status"`
	Book          Book         `json:"book"`
	User          User         `json:"user"`
	Admin_id      int          `json:"admin_id"`
	Penalties     Penalty      `json:"penalties"`
	Created_at    time.Time    `json:"created_at"`
	Updated_at    time.Time    `json:"updated_at"`
}

// type DetailBookLoanResponse struct {
// 	Loan_id       int          `json:"loan_id"`
// 	Checkout_date time.Time    `json:"checkout_date"`
// 	Due_date      time.Time    `json:"due_date"`
// 	Return_date   sql.NullTime `json:"return_date"`
// 	Status        string       `json:"status"`
// 	Admin_id      int          `json:"admin_id"`
// 	Book          Book         `json:"book"`
// 	User          User         `json:"user"`
// 	Penalty       Penalty      `json:"penalty"`
// 	Created_at    time.Time    `json:"created_at"`
// 	Updated_at    time.Time    `json:"updated_at"`
// }
