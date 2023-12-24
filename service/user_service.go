package service

import (
	"context"

	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
)

type UserService interface {
	CreateUser(ctx context.Context, request webrequest.UserCreateRequest ) webresponse.UserResponse
	Login(ctx context.Context, request webrequest.UserLoginRequest )webresponse.LoginResponse
	Authenticate(ctx context.Context, id int) webresponse.UserResponse
}

