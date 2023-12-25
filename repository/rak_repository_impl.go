package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
)

type RakRepositoryImpl struct {
}

func NewRakRepository() RakRepository {
	return &RakRepositoryImpl{}
}

func (r *RakRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, rak domain.Rak) domain.Rak {
	SQL := "insert into rak(name,col,rows_rak) values(?,?,?)"

	result, err := tx.ExecContext(ctx, SQL, rak.Name, rak.Col, rak.Rows_rak)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	rak.Rak_id = int(id)
	return rak
}

func (r *RakRepositoryImpl) FindByNameColRow(ctx context.Context, tx *sql.Tx, rak webrequest.RakByNameRowRequest) (domain.Rak, error) {
	SQL := "select rak_id, name, col, rows_rak,created_at, updated_at from rak where name = ? and rows_rak = ? and col = ?"

	rows, err := tx.QueryContext(ctx, SQL, rak.Name, rak.Rows_rak, rak.Rows_rak)
	helper.PanicIfError(err)
	defer rows.Close()

	ra := domain.Rak{}

	if rows.Next() {
		err := rows.Scan(&ra.Rak_id, &ra.Name, &ra.Col, &ra.Rows_rak, &ra.Created_at, &ra.Updated_at)
		helper.PanicIfError(err)
		return ra, nil
	} else {
		return ra, errors.New("rak not found")
	}
}
