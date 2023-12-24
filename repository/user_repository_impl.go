package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
)

type UserRepositoryImpl struct {

}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
} 

func (r *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into user(name,email,password,level,is_enabled,gender,telp,birthdate,address,foto,batas) values(?,?,?,?,?,?,?,?,?,?,?)"

	result, err:= tx.ExecContext(ctx, SQL, user.Name,user.Email,user.Password, user.Level,user.Is_enabled,user.Gender,user.Telp,user.Birthdate,user.Address,user.Foto,user.Batas)
	helper.PanicIfError(err)
	
	fmt.Println(result)
	id,err:=result.LastInsertId()

	user.User_id = int(id)
	return user

}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "select user_id, name,email, password, level,is_enabled,gender, telp, birthdate, address, foto, batas,created_at,updated_at,deleted_at from user where email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()
	fmt.Println("repo findbyemail")
	// fmt.Println(rows)

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.User_id,&user.Email,&user.Password,&user.Level,&user.Gender, &user.Telp, &user.Birthdate, &user.Address, &user.Foto, &user.Batas,&user.Created_at,&user.Updated_at,&user.Deleted_at)
		fmt.Println(err)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("email not found")
	}

}
