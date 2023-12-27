package service

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

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
	fmt.Println("service jalan")
	// service
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUser, err := s.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err == nil {
		fmt.Println("email sama " + getUser.Email)
		// panic("email sama")
		// err34 = errors.New("email sama")

		panic(exception.DuplicateEmailError{Error: "email already exists"})
	}

	reader := bytes.NewReader(request.Foto)

	result, err := s.Cld.Upload.Upload(ctx, reader, uploader.UploadParams{})
	if err != nil {
		fmt.Println(err)
		panic("upload fatal")
	}
	fmt.Println(result.SecureURL)

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		fmt.Println(err)
		panic("failed hashing password")
	}
	fmt.Println(hashedPassword)

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
	fmt.Println("service jalan")
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// find user by email
	getUser, err := s.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "email " + request.Email + " not found"})
	}
	fmt.Println(getUser)

	// compare password
	comparePassword := helper.ComparePassword(request.Password, getUser.Password)

	if !comparePassword {
		panic(exception.CustomEror{Code: 400, Error: "password not match "})
	}
	fmt.Println(comparePassword)

	toString := webrequest.UserGenereteToken{
		Id:    getUser.User_id,
		Email: getUser.Email,
		Level: getUser.Level,
	}
	generateToken, err := helper.GenerateJWT(toString)
	helper.PanicIfError(err)
	fmt.Println(generateToken)

	token := webresponse.LoginResponse{
		Token: generateToken,
	}
	return token
}

func (s *UserServiceImpl) Authenticate(ctx context.Context, id int) webresponse.UserResponse {
	fmt.Println("service jalan")

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUser, err := s.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user not found"})
	}
	fmt.Println(getUser)
	// user := domain.User{
	// 	Name:       getUser.Name,
	// 	Email:      getUser.Email,
	// 	Level:      getUser.Level,
	// 	Password:   getUser.Password,
	// 	Is_enabled: getUser.Is_enabled,
	// 	Gender:     getUser.Gender,
	// 	Telp:       getUser.Telp,
	// 	Birthdate:  getUser.Birthdate,
	// 	Address:    getUser.Address,
	// 	Foto:       getUser.Foto,
	// 	Batas:      getUser.Batas,
	// }

	return helper.ToUserResponse(getUser)
}

func (s *UserServiceImpl) ListAllUsers(ctx context.Context) []webresponse.UserResponse {
	fmt.Println("Listing all users")

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUsers := s.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(getUsers)

}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, request webrequest.UpdateUserRequest, id int) bool {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = s.UserRepository.FindById(ctx, tx, id)
	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.Foto != nil {
		reader := bytes.NewReader(request.Foto)

		result, err := s.Cld.Upload.Upload(ctx, reader, uploader.UploadParams{})
		if err != nil {
			fmt.Println(err)
			panic("upload fatal")
		}
		fmt.Println(result.SecureURL)
		request.UrlFoto = result.SecureURL
	}

	update := s.UserRepository.Update(ctx, tx, id, request)
	fmt.Println(update)
	// user := domain.User{}
	// panic("sda")
	return true
}

func (s *UserServiceImpl) FindByid(ctx context.Context, id string) webresponse.UserResponse {
	fmt.Println("find by id userservice ")
	idInt := konversi.StrToInt(id, "user id")

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUser, err := s.UserRepository.FindById(ctx, tx, idInt)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user not found"})
	}
	fmt.Println(getUser)

	return helper.ToUserResponse(getUser)

}
