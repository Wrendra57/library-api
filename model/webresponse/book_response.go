package webresponse

import (
	"time"

	"github.com/be/perpustakaan/model/domain"
)

type BookResponse struct {
	Book_id          int              `json:"book_id"`
	Title            string           `json:"title"`
	Category         domain.Category  `json:"category"`
	Author           domain.Author    `json:"author"`
	Publisher        domain.Publisher `json:"publisher"`
	Isbn             string           `json:"isbn"`
	Page_count       int              `json:"page_count"`
	Stock            int              `json:"stock"`
	Publication_year int              `json:"publication"`
	Foto             string           `json:"foto"`
	Rak_id           domain.Rak       `json:"rak"`
	Price            int              `json:"price"`
	Created_at       time.Time        `json:"created_at"`
	Updated_at       time.Time        `json:"updated_at"`
}
