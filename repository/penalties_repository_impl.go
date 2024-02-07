package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
)

type PenaltiesRepositoryImpl struct {
}

func NewPenaltiesRepository() PenaltiesRepository {
	return &PenaltiesRepositoryImpl{}
}

func (r *PenaltiesRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, p domain.Penalties) domain.Penalties {
	SQL := ` INSERT INTO penalties(loan_id,penalty_amount,reason,payment_status,due_date,admin_id) VALUES(?,?,?,?,?,?)`

	result, err := tx.ExecContext(ctx, SQL, p.Loan_id, p.Penalty_amount, p.Reason, p.Payment_status, p.Due_date, p.Admin_id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	p.Penalty_id = int(id)
	p.Created_at = time.Now()
	p.Updated_at = time.Now()
	return p
}

func (r *PenaltiesRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Penalties, error) {
	SQL := `SELECT penalty_id,loan_id,penalty_amount,reason,payment_status,payment_status,due_date,admin_id,created_at,updated_at from penalties where loan_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)

	defer rows.Close()
	p := domain.Penalties{}
	if rows.Next() {
		err := rows.Scan(&p.Penalty_id, &p.Loan_id, &p.Penalty_amount, &p.Reason, &p.Payment_status, &p.Due_date, &p.Admin_id, &p.Created_at, &p.Updated_at)
		helper.PanicIfError(err)
		return p, nil
	} else {
		return p, errors.New("book loan not found")
	}
}
