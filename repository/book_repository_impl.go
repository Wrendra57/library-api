package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
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
				b.updated_at,
				b.deleted_at
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
		err := rows.Scan(&b.Book_id, &b.Title, &b.Category, &b.Author, &b.Publisher, &b.Isbn, &b.Page_count, &b.Stock, &b.Publication_year, &b.Foto, &b.Rak, &b.Column, &b.Rows_rak, &b.Price, &b.Admin, &b.Created_at, &b.Updated_at, &b.Deleted_at)

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
				b.updated_at,
				b.deleted_at
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

		err := rows.Scan(&b.Book_id, &b.Title, &b.Category, &b.Author, &b.Publisher, &b.Isbn, &b.Page_count, &b.Stock, &b.Publication_year, &b.Foto, &b.Rak, &b.Column, &b.Rows_rak, &b.Price, &b.Admin, &b.Created_at, &b.Updated_at, &b.Deleted_at)
		helper.PanicIfError(err)
		books = append(books, b)
	}
	return books
}

func (r *BookRepositoryImpl) FindBook(ctx context.Context, tx *sql.Tx, s webrequest.SearchBookRequest) []webresponse.BookResponseComplete {
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
				b.updated_at,
				b.deleted_at
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
				OR LOWER(b.title) LIKE LOWER(CONCAT('%', ?, '%'))
				OR LOWER(c.category) LIKE LOWER(CONCAT('%', ?, '%'))
				OR LOWER(a.name) LIKE LOWER(CONCAT('%', ?, '%'))
				OR LOWER(p.name) LIKE LOWER(CONCAT('%', ?, '%'))
				OR LOWER(b.isbn) LIKE LOWER(CONCAT('%', ?, '%'))
			limit ? 
			offset  ?
	`
	rows, err := tx.QueryContext(ctx, SQL, s.Search, s.Search, s.Search, s.Search, s.Search, s.Search, s.Limit, s.Offset)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []webresponse.BookResponseComplete
	for rows.Next() {
		b := webresponse.BookResponseComplete{}

		err := rows.Scan(&b.Book_id, &b.Title, &b.Category, &b.Author, &b.Publisher, &b.Isbn, &b.Page_count, &b.Stock, &b.Publication_year, &b.Foto, &b.Rak, &b.Column, &b.Rows_rak, &b.Price, &b.Admin, &b.Created_at, &b.Updated_at, &b.Deleted_at)
		helper.PanicIfError(err)
		books = append(books, b)
	}
	return books

}

func (r *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, id int, book domain.Book) int {
	SQL := "update book set "
	var args []interface{}

	if book.Title != "" {
		SQL += "title = ?, "
		args = append(args, book.Title)
	}
	if book.Category_id != 0 {
		SQL += "category_id = ?, "
		args = append(args, book.Category_id)
	}
	if book.Author_id != 0 {
		SQL += "author_id = ?, "
		args = append(args, book.Author_id)
	}
	if book.Publisher_id != 0 {
		SQL += "publisher_id = ?, "
		args = append(args, book.Publisher_id)
	}
	if book.Isbn != "" {
		SQL += "isbn = ?, "
		args = append(args, book.Isbn)
	}
	if book.Page_count != 0 {
		SQL += "page_count = ?, "
		args = append(args, book.Page_count)
	}
	if book.Stock != 0 {
		SQL += "stock = ?, "
		args = append(args, book.Stock)
	}
	if book.Publication_year != 0 {
		SQL += "publication_year = ?, "
		args = append(args, book.Publication_year)
	}
	if book.Foto != "" {
		SQL += "foto = ?, "
		args = append(args, book.Foto)
	}
	if book.Rak_id != 0 {
		SQL += "rak_id = ?, "
		args = append(args, book.Rak_id)
	}
	if book.Price != 0 {
		SQL += "price = ?, "
		args = append(args, book.Price)
	}
	if book.Admin_id != 0 {
		SQL += "admin_id = ?, "
		args = append(args, book.Admin_id)
	}
	SQL = SQL[:len(SQL)-2]

	SQL += " WHERE book_id = ?"
	args = append(args, id)
	_, err := tx.Exec(SQL, args...)
	helper.PanicIfError(err)

	return id
}

func (r *BookRepositoryImpl) DeleteBook(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := "update book set deleted_at = ?, admin_id = ? where book_id = ?"
	_, err := tx.ExecContext(ctx, SQL, time.Now(), ctx.Value("id").(int), id)
	helper.PanicIfError(err)
	if err != nil {
		return err
	}
	return nil
}
