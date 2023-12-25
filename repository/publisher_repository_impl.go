package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
)

type PublisherRepositoryImpl struct {
}

func NewPublisherRepository() PublisherRepository {
	return &PublisherRepositoryImpl{}
}

func (r *PublisherRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, publisher domain.Publisher) domain.Publisher {
	SQL := "insert into publisher(name) values(?)"

	result, err := tx.ExecContext(ctx, SQL, publisher.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	publisher.Publisher_id = int(id)
	return publisher

}

func (r *PublisherRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Publisher, error) {
	SQL := "select publisher_id, name, created_at, updated_at from publisher where publisher_id = ?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	publisher := domain.Publisher{}
	if rows.Next() {
		err := rows.Scan(&publisher.Publisher_id, &publisher.Name, &publisher.Created_at, &publisher.Updated_at)
		helper.PanicIfError(err)
		return publisher, nil
	} else {
		return publisher, errors.New("publisher not founf")
	}

}

func (r *PublisherRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Publisher, error) {
	SQL := "select publisher_id,name,created_at,updated_at from publisher where name=?"
	rows, err := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(err)
	defer rows.Close()

	publisher := domain.Publisher{}
	if rows.Next() {
		err := rows.Scan(&publisher.Publisher_id, &publisher.Name, &publisher.Created_at, &publisher.Updated_at)
		helper.PanicIfError(err)
		return publisher, nil
	} else {
		return publisher, errors.New("publiher not founf")
	}
}

func (r *PublisherRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Publisher {
	SQL := "select publisher_id, name, created_at, updated_at from publisher"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var publisher []domain.Publisher
	for rows.Next() {
		a := domain.Publisher{}
		err := rows.Scan(&a.Publisher_id, &a.Name, &a.Created_at, &a.Updated_at)
		helper.PanicIfError(err)
		publisher = append(publisher, a)
	}
	return publisher
}
