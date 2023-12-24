package service

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	// "github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
	"github.com/be/perpustakaan/repository"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository 	repository.UserRepository
	DB				*sql.DB
	Validate		*validator.Validate
	Cld	*cloudinary.Cloudinary		  		
}
func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate, cld	*cloudinary.Cloudinary) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: validate,
		Cld: cld,
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, request webrequest.UserCreateRequest ) webresponse.UserResponse{
	fmt.Println("service jalan")
	// service
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUser,err:= s.UserRepository.FindByEmail(ctx, tx,request.Email)
	if err == nil {
		fmt.Println("email sama " + getUser.Email)
		// panic("email sama")
		// err34 = errors.New("email sama")
		panic(exception.DuplicateEmailError{Error:"email already exists"})
	}

	reader :=bytes.NewReader(request.Foto)

	result,err:=s.Cld.Upload.Upload(ctx,reader,uploader.UploadParams{})
	if err != nil {
		fmt.Println(err)
		panic("upload fatal")
    }
	fmt.Println(result.SecureURL)

	hashedPassword,err:=helper.HashPassword(request.Password)
	if err != nil {
		fmt.Println(err)
		panic("failed hashing password")
	}
	fmt.Println(hashedPassword)

	user := domain.User{
	Name: request.Name,
	Email: request.Email,
	Password: hashedPassword,
	Level: "member",
	Is_enabled: true,
	Gender: request.Gender,
	Telp : request.Telp,
	Birthdate :request.Birthdate,
	Address : request.Address,
	Foto : result.SecureURL,
	Batas :3,
	}

	user = s.UserRepository.Create(ctx,tx,user)

	return helper.ToUserResponse(user)
}


