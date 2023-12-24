package repository

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
}