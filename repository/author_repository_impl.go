package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
)

type AuthorRepositoryImpl struct {
}

func NewAuthorRepository() AuthorRepository {
	return &AuthorRepositoryImpl{}
}

func (r *AuthorRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, author domain.Author) domain.Author {
	SQL := "insert into author(name) values(?)"

	result, err := tx.ExecContext(ctx, SQL, author.Name)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	author.Author_id = int(id)
	return author
	// panic("das")
}

func (r *AuthorRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Author, error) {
	SQL := "select author_id,name,created_at,updated_at from author where author_id=?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	author := domain.Author{}
	if rows.Next() {
		err := rows.Scan(&author.Author_id, &author.Name, &author.Created_at, &author.Updated_at)
		helper.PanicIfError(err)
		return author, nil

	} else {
		return author, errors.New("author not found")
	}
	// panic("sda")
}

func (r *AuthorRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Author, error) {

	SQL := "select author_id, name, created_at, updated_at from author where name=?"

	rows, err := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(err)
	defer rows.Close()

	var a domain.Author

	if rows.Next() {
		err := rows.Scan(&a.Author_id, &a.Name, &a.Created_at, &a.Updated_at)
		helper.PanicIfError(err)
		return a, nil
	} else {
		return a, errors.New("author not found")
	}
}

func (r *AuthorRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Author {
	SQL := "select author_id, name, created_at, updated_at from author"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var authors []domain.Author
	for rows.Next() {
		a := domain.Author{}
		err := rows.Scan(&a.Author_id, &a.Name, &a.Created_at, &a.Updated_at)
		helper.PanicIfError(err)
		authors = append(authors, a)
	}
	return authors

}
