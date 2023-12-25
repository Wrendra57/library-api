package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx)[]domain.User
	Update(ctx context.Context, tx *sql.Tx, id int, user webrequest.UpdateUserRequest) webrequest.UpdateUserRequest
}