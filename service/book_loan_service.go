package service

import (
	"context"

	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
)

type BookLoanService interface {
	CreateBookLoan(ctx context.Context, request webrequest.BookLoanCreateRequest) webresponse.BookLoanResponseComplete
}
