package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
)

type BookRepositoryImpl struct{}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (r *BookRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, b domain.Book) domain.Book {
	SQL := `insert into book(title, category_id,author_id,publisher_id,isbn,page_count,stock,publication_year,foto,rak_id,price,admin_id) values(?,?,?,?,?,?,?,?,?,?,?,?)`

	result, err := tx.ExecContext(ctx, SQL, b.Title, b.Category_id, b.Author_id, b.Publisher_id, b.Isbn, b.Page_count, b.Stock, b.Publication_year, b.Foto, b.Rak_id, b.Price, b.Admin_id)

	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	b.Book_id = int(id)
	b.Created_at = time.Now()
	b.Updated_at = time.Now()
	return b

}
