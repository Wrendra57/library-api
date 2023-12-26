package webrequest

import "time"

type BookCreateRequest struct {
	Title            string `validate:"required,min=3,max=255" json:"title"`
	Category         string `validate:"required,min=2,max=255" json:"category"`
	Author           string `validate:"required,min=2,max=255" json:"author"`
	Publisher        string `validate:"required,min=2,max=255" json:"publisher"`
	Isbn             string `validate:"required,min=2,max=255" json:"isbn"`
	Page_count       string `validate:"required" json:"page_count"`
	Stock            string `validate:"required" json:"stock"`
	Publication_year string `validate:"required" json:"publication_year"`
	Foto             []byte `validate:"required" json:"foto"`
	Rak              string `validate:"required" json:"rak"`
	Column           string `validate:"required" json:"column"`
	Rows             string `validate:"required" json:"rows"`
	Price            string `validate:"required" json:"price"`
}

type UserCsreateRequests struct {
	Name      string    `validate:"required,min=3,max=100" json:"name"`
	Email     string    `validate:"required,min=1,max=100,email" json:"email"`
	Password  string    `validate:"required,min=4,max=100" json:"password"`
	Gender    string    `validate:"required" json:"gender"`
	Telp      string    `validate:"required" json:"telp"`
	Birthdate time.Time `validate:"required" json:"birthdate"`
	Address   string    `validate:"required" json:"address"`
	Foto      []byte    `validate:"required" json:"foto"`
}

type FindAllRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

type SearchBookRequest struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
