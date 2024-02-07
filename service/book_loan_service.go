package service

import (
	"context"

	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
)

type BookLoanService interface {
	CreateBookLoan(ctx context.Context, request webrequest.BookLoanCreateRequest) webresponse.BookLoanResponseComplete
	ReturnBookLoan(ctx context.Context, request webrequest.BookLoanCreateRequest) webresponse.BookLoanResponseComplete2
	FindAll(ctx context.Context, request webrequest.ListALlBookLoanRequest) []webresponse.ListBookLoanResponse
	FindById(ctx context.Context, id int) webresponse.ListBookLoanResponse
}
