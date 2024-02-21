package service

import (
	"bytes"
	"context"
	"database/sql"

	// "github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/helper/konversi"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
	"github.com/be/perpustakaan/repository"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
	Cld            *cloudinary.Cloudinary
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate, cld *cloudinary.Cloudinary) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
		Cld:            cld,
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, request webrequest.UserCreateRequest) webresponse.UserResponse {

	// service
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = s.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err == nil {
		panic(exception.DuplicateEmailError{Error: "Email already exists"})
	}

	reader := bytes.NewReader(request.Foto)

	result, err := s.Cld.Upload.Upload(ctx, reader, uploader.UploadParams{})
	if err != nil {

		panic(exception.DuplicateEmailError{Error: "upload fatal"})
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		panic(exception.DuplicateEmailError{Error: "failed hashing password"})

	}

	user := domain.User{
		Name:       request.Name,
		Email:      request.Email,
		Password:   hashedPassword,
		Level:      "member",
		Is_enabled: true,
		Gender:     request.Gender,
		Telp:       request.Telp,
		Birthdate:  request.Birthdate,
		Address:    request.Address,
		Foto:       result.SecureURL,
		Batas:      3,
	}

	user = s.UserRepository.Create(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (s *UserServiceImpl) Login(ctx context.Context, request webrequest.UserLoginRequest) webresponse.LoginResponse {

	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// find user by email
	getUser, err := s.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "Email " + request.Email + " not found"})
	}

	// compare password
	comparePassword := helper.ComparePassword(request.Password, getUser.Password)

	if !comparePassword {
		panic(exception.CustomEror{Code: 400, Error: "Password not match"})
	}

	toString := webrequest.UserGenereteToken{
		Id:    getUser.User_id,
		Email: getUser.Email,
		Level: getUser.Level,
	}
	generateToken, err := helper.GenerateJWT(toString)
	helper.PanicIfError(err)

	token := webresponse.LoginResponse{
		Token: generateToken,
	}
	return token
}

func (s *UserServiceImpl) Authenticate(ctx context.Context, id int) webresponse.UserResponse {

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUser, err := s.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user not found"})
	}

	return helper.ToUserResponse(getUser)
}

func (s *UserServiceImpl) ListAllUsers(ctx context.Context) []webresponse.UserResponse {

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUsers := s.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(getUsers)

}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, request webrequest.UpdateUserRequest, id int) bool {
	level, ok := ctx.Value("level").(string)
	idToken, _ := ctx.Value("id").(int)
	if !ok {
		panic(exception.CustomEror{Code: 400, Error: "token not found "})
	}

	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.FindById(ctx, tx, id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if level != "superadmin" && user.User_id != idToken {
		panic(exception.CustomEror{Code: 400, Error: "unautorized access"})
	}
	if level != "superadmin" {
		request.Level = ""
	}
	if request.Foto != nil {
		reader := bytes.NewReader(request.Foto)

		result, err := s.Cld.Upload.Upload(ctx, reader, uploader.UploadParams{})
		if err != nil {

			panic("upload fatal")
		}

		request.UrlFoto = result.SecureURL
	}

	_ = s.UserRepository.Update(ctx, tx, id, request)

	return true
}

func (s *UserServiceImpl) FindByid(ctx context.Context, id string) webresponse.UserResponse {

	idInt := konversi.StrToInt(id, "user id")

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUser, err := s.UserRepository.FindById(ctx, tx, idInt)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user not found"})
	}

	return helper.ToUserResponse(getUser)

}
