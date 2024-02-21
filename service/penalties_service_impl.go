package service

import (
	"context"
	"database/sql"

	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/repository"
	"github.com/go-playground/validator/v10"
)

type PenaltiesServiceImpl struct {
	UserRepository      repository.UserRepository
	PenaltiesRepository repository.PenaltiesRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewPenaltiesService(userRepository repository.UserRepository, penaltiesRepository repository.PenaltiesRepository, DB *sql.DB, validate *validator.Validate) PenaltiesService {
	return &PenaltiesServiceImpl{
		UserRepository:      userRepository,
		PenaltiesRepository: penaltiesRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (s *PenaltiesServiceImpl) PayPenalties(ctx context.Context, r webrequest.UpdatePenaltiesRequest, id int) domain.Penalties {
	admin_id, ok := ctx.Value("id").(int)

	if !ok {
		panic(exception.CustomEror{Code: 400, Error: "user not found"})
	}

	err := s.Validate.Struct(r)
	helper.PanicIfError(err)

	p := domain.Penalties{}

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.FindById(ctx, tx, admin_id)
	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "user unauthorized"})
	}
	p, err = s.PenaltiesRepository.FindById(ctx, tx, id)

	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "Penalties not found"})
	}
	if p.Payment_status == "paid" {
		panic(exception.CustomEror{Code: 400, Error: "Penalties was pay"})
	}
	if p.Penalty_amount != r.Penalty_amount {
		panic(exception.CustomEror{Code: 400, Error: "Amount must same with penalty amount"})
	}

	r.Admin_id = user.User_id

	update := s.PenaltiesRepository.Update(ctx, tx, id, r)

	p.Admin_id = update.Admin_id
	p.Payment_status = update.Payment_status

	return p
}
