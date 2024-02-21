package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into user(name,email,password,level,is_enabled,gender,telp,birthdate,address,foto,batas) values(?,?,?,?,?,?,?,?,?,?,?)"

	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Level, user.Is_enabled, user.Gender, user.Telp, user.Birthdate, user.Address, user.Foto, user.Batas)
	helper.PanicIfError(err)

	// fmt.Println(result)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.User_id = int(id)
	return user

}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "select user_id, name,email, password, level,is_enabled,gender, telp, birthdate, address, foto, batas,created_at,updated_at,deleted_at from user where email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()
	// fmt.Println("repo findbyemail")
	// fmt.Println(rows)

	user := domain.User{}
	if rows.Next() {

		err := rows.Scan(&user.User_id, &user.Name, &user.Email, &user.Password, &user.Level, &user.Is_enabled, &user.Gender, &user.Telp, &user.Birthdate, &user.Address, &user.Foto, &user.Batas, &user.Created_at, &user.Updated_at, &user.Deleted_at)
		// fmt.Println(err)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}

}

func (r *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	SQL := "select user_id, name,email, password, level,is_enabled,gender, telp, birthdate, address, foto, batas,created_at,updated_at,deleted_at from user where user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Name, &user.Email, &user.Password, &user.Level, &user.Is_enabled, &user.Gender, &user.Telp, &user.Birthdate, &user.Address, &user.Foto, &user.Batas, &user.Created_at, &user.Updated_at, &user.Deleted_at)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select user_id, name,email, password, level,is_enabled,gender, telp, birthdate, address, foto, batas,created_at,updated_at,deleted_at from user"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.User_id, &user.Name, &user.Email, &user.Password, &user.Level, &user.Is_enabled, &user.Gender, &user.Telp, &user.Birthdate, &user.Address, &user.Foto, &user.Batas, &user.Created_at, &user.Updated_at, &user.Deleted_at)
		helper.PanicIfError(err)
		users = append(users, user)

	}
	return users

	// panic("sad")
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, id int, user webrequest.UpdateUserRequest) webrequest.UpdateUserRequest {

	SQL := "update user set "
	var args []interface{}

	if user.Name != "" {
		SQL += "name = ?, "
		args = append(args, user.Name)
	}
	if user.Email != "" {
		SQL += "email = ?, "
		args = append(args, user.Email)
	}
	if user.Gender != "" {
		SQL += "gender = ?, "
		args = append(args, user.Gender)
	}
	if user.Telp != "" {
		SQL += "telp = ?, "
		args = append(args, user.Telp)
	}
	if user.Birthdate.Valid {
		SQL += "birthdate = ?, "
		args = append(args, user.Birthdate.Time)
	}
	if user.UrlFoto != "" {
		SQL += "foto = ?, "
		args = append(args, user.UrlFoto)
	}
	if user.Level != "" {
		SQL += "level = ?, "
		args = append(args, user.Level)
	}
	if user.Is_enabled != "" {
		SQL += "is_enabled = ?, "

		if user.Is_enabled == "true" {
			args = append(args, true)
		} else {
			args = append(args, false)
		}
	}
	if user.Batas != "" {
		SQL += "batas = ?, "
		number, err := strconv.Atoi(user.Batas)
		helper.PanicIfError(err)
		args = append(args, number)
	}
	SQL = SQL[:len(SQL)-2]
	SQL += " WHERE user_id = ?"
	args = append(args, id)

	fmt.Println(SQL)
	_, err := tx.Exec(SQL, args...)
	if err != nil {
		panic(err)
	}
	// fmt.Println(result)

	return user
}
