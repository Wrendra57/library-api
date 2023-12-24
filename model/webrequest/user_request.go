package webrequest

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserCreateRequest struct {
	Name string `validate:"required,min=3,max=100" json:"name"`
	Email string  `validate:"required,min=1,max=100,email" json:"email"`
	Password string `validate:"required,min=4,max=100" json:"password"`
	Gender string `validate:"required" json:"gender"`
	Telp string `validate:"required" json:"telp"`
	Birthdate time.Time `validate:"required" json:"birthdate"`
	Address string `validate:"required" json:"address"`
	Foto []byte `validate:"required" json:"foto"`
}

type UserLoginRequest struct {
	Email string `validate:"required,min=1,max=100,email" json:"email"`
	Password string `validate:"required,min=4,max=100" json:"password"`
}

type UserGenereteToken struct {
	
	Id int	`json:"id"`
	Email string `json:"email"`
	Level string `json:"level"`
	jwt.RegisteredClaims
}