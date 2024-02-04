package helper

import (
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webresponse"
)

func ToUserResponse(user domain.User) webresponse.UserResponse {
	return webresponse.UserResponse{
		User_id:    user.User_id,
		Name:       user.Name,
		Email:      user.Email,
		Level:      user.Level,
		Is_enabled: user.Is_enabled,
		Gender:     user.Gender,
		Telp:       user.Telp,
		Birthdate:  user.Birthdate,
		Address:    user.Address,
		Foto:       user.Foto,
		Batas:      user.Batas,
	}
}

func ToUserResponses(users []domain.User) []webresponse.UserResponse {
	var userResponses []webresponse.UserResponse

	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToBookLoanResponse(loan domain.BookLoan) webresponse.BookLoanResponseComplete {
	return webresponse.BookLoanResponseComplete{
		Loan_id:       loan.Loan_id,
		Checkout_date: loan.Checkout_date,
		Due_date:      loan.Due_date,
		Return_date:   loan.Return_date,
		Status:        loan.Status,
		Book_id:       loan.Book_id,
		User_id:       loan.User_id,
		Admin_id:      loan.Admin_id,
		Created_at:    loan.Created_at,
		Updated_at:    loan.Updated_at,
	}
}
