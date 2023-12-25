package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
)

type CategoryBookRepositoryImpl struct {
}

func NewCategoryRepository() CategoryBookRepository {
	return &CategoryBookRepositoryImpl{}
}

func (r *CategoryBookRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(category) values(?)"

	result, err := tx.ExecContext(ctx, SQL, category.Category)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Category_id = int(id)
	return category
	// panic("ds")
}

func (r *CategoryBookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	SQL := "select category_id,category,created_at,updated_at from category where category_id=?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Category_id, &category.Category, &category.Created_at, &category.Updated_at)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not founf")
	}
}

func (r *CategoryBookRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, category string) (domain.Category, error) {
	SQL := "select category_id,category,created_at,updated_at from category where category=?"

	rows, err := tx.QueryContext(ctx, SQL, category)
	helper.PanicIfError(err)
	defer rows.Close()

	ctr := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&ctr.Category_id, &ctr.Category, &ctr.Created_at, &ctr.Updated_at)
		helper.PanicIfError(err)
		return ctr, nil
	} else {
		return ctr, errors.New("category not found")
	}

}

func (r *CategoryBookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select category_id, category, created_at, updated_at from category"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var category []domain.Category
	for rows.Next() {
		a := domain.Category{}
		err := rows.Scan(&a.Category_id, &a.Category, &a.Created_at, &a.Updated_at)
		helper.PanicIfError(err)
		category = append(category, a)
	}
	return category

}
