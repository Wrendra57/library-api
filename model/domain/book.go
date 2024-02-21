package domain

import (
	"database/sql"
	"time"
)

type Book struct {
	Book_id          int
	Title            string
	Category_id      int
	Author_id        int
	Publisher_id     int
	Isbn             string
	Page_count       int
	Stock            int
	Publication_year int
	Foto             string
	Rak_id           int
	Price            int
	Admin_id         int
	Created_at       time.Time
	Updated_at       time.Time
	Deleted_at       sql.NullTime
	UpdateForRent    string
}
