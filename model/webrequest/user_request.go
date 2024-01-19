package webrequest

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type StatusEnum string

const (
	MemberLevel     StatusEnum = "member"
	AdminLevel      StatusEnum = "admin"
	SuperAdminLevel StatusEnum = "superadmin"
)

type GenderEnum string

const (
	Male   GenderEnum = "male"
	Female GenderEnum = "female"
)

type UserCreateRequest struct {
	Name      string    `validate:"required,min=3,max=100" json:"name"`
	Email     string    `validate:"required,min=1,max=100,email" json:"email"`
	Password  string    `validate:"required,min=1,max=100" json:"password"`
	Gender    string    `validate:"required,oneof=male female" json:"gender"`
	Telp      string    `validate:"required" json:"telp"`
	Birthdate time.Time `validate:"required" json:"birthdate"`
	Address   string    `validate:"required" json:"address"`
	Foto      []byte    `validate:"required" json:"foto"`
}

type UserLoginRequest struct {
	Email    string `validate:"required,min=1,max=100,email" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}

type UserGenereteToken struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Level string `json:"level"`
	jwt.RegisteredClaims
}

type UpdateUserRequest struct {
	Name       string       `validate:"omitempty,min=3,max=100" json:"name"`
	Email      string       `validate:"omitempty,min=3,max=100,email" json:"email"`
	Gender     GenderEnum   `validate:"omitempty,oneof=male female" json:"gender"`
	Telp       string       `validate:"omitempty,min=8,max=14" json:"telp"`
	Birthdate  sql.NullTime `json:"birthdate"`
	Address    string       `validate:"omitempty,min=3,max=255" json:"address"`
	Foto       []byte       `json:"foto"`
	Level      StatusEnum   `validate:"omitempty,oneof=member admin superadmin" json:"level"`
	Is_enabled string       `json:"is_enabled"`
	Batas      string       `validate:"omitempty" json:"batas"`
	UrlFoto    string       `json:"url_foto"`
}

// type UpdateUserRequest2 struct {
// 	Name sql.NullString `json:"name"`
// 	Email sql.NullString  `json:"email"`
// 	Gender sql.NullString ` json:"gender"`
// 	Telp sql.NullString `json:"telp"`
// 	Birthdate sql.NullTime `json:"birthdate"`
// 	Address sql.NullString `json:"address"`
// 	Foto sql.NullByte `json:"foto"`
// 	Level sql.NullTime `json:"level"`
// 	Is_enabled sql.NullBool `json:"is_enabled"`
// 	Batas sql.NullInt64 	`json:"batas"`
// }
