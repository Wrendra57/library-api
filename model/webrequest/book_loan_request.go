package webrequest

type BookLoanCreateRequest struct {
	Book_id int `validate:"required,number" json:"book_id"`
	User_id int `validate:"required,number" json:"user_id"`
}
