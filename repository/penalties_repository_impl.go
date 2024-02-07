package repository

import (
	"context"
	"database/sql"
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
