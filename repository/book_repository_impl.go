package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webresponse"
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

func (r *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (webresponse.BookResponseComplete, error) {
	SQL := `
			SELECT
				b.book_id,
				b.title,
				c.category,
				a.name AS author,
				p.name AS publisher,
				b.isbn,
				b.page_count,
				b.stock,
				b.publication_year,
				b.foto,
				r.name AS rak,
				r.col,
				r.rows_rak,
				b.price,
				u.name AS admin,
				b.created_at,
				b.updated_at
			FROM
				book b
			JOIN
				category c ON b.category_id = c.category_id
			JOIN
				author a ON b.author_id = a.author_id
			JOIN
				publisher p ON b.publisher_id = p.publisher_id
			JOIN
				rak r ON b.rak_id = r.rak_id
			JOIN
				user u ON b.admin_id = u.user_id
			WHERE
				b.book_id = ?
	`

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	b := webresponse.BookResponseComplete{}
	if rows.Next() {
		err := rows.Scan(&b.Book_id, &b.Title, &b.Category, &b.Author, &b.Publisher, &b.Isbn, &b.Page_count, &b.Stock, &b.Publication_year, &b.Foto, &b.Rak, &b.Column, &b.Rows_rak, &b.Price, &b.Admin, &b.Created_at, &b.Updated_at)

		helper.PanicIfError(err)
		return b, nil
	} else {
		return b, errors.New("user not found")
	}
}

func (r *BookRepositoryImpl) ListBook(ctx context.Context, tx *sql.Tx, limit int, offset int) []webresponse.BookResponseComplete {
	SQL := `
			SELECT
				b.book_id,
				b.title,
				c.category,
				a.name AS author,
				p.name AS publisher,
				b.isbn,
				b.page_count,
				b.stock,
				b.publication_year,
				b.foto,
				r.name AS rak,
				r.col,
				r.rows_rak,
				b.price,
				u.name AS admin,
				b.created_at,
				b.updated_at
			FROM
				book b
			JOIN
				category c ON b.category_id = c.category_id
			JOIN
				author a ON b.author_id = a.author_id
			JOIN
				publisher p ON b.publisher_id = p.publisher_id
			JOIN
				rak r ON b.rak_id = r.rak_id
			JOIN
				user u ON b.admin_id = u.user_id
			limit ? 
			offset  ?
	`

	rows, err := tx.QueryContext(ctx, SQL, limit, offset)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []webresponse.BookResponseComplete
	for rows.Next() {
		b := webresponse.BookResponseComplete{}

		err := rows.Scan(&b.Book_id, &b.Title, &b.Category, &b.Author, &b.Publisher, &b.Isbn, &b.Page_count, &b.Stock, &b.Publication_year, &b.Foto, &b.Rak, &b.Column, &b.Rows_rak, &b.Price, &b.Admin, &b.Created_at, &b.Updated_at)
		helper.PanicIfError(err)
		books = append(books, b)
	}
	return books
}
