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

func ToBookLoanResponseComplete(l domain.BookLoan, b webresponse.BookResponseComplete, u domain.User, a domain.User, p domain.Penalties) webresponse.BookLoanResponseComplete2 {
	user := ToUserResponse(u)
	admin := ToUserResponse(a)
	return webresponse.BookLoanResponseComplete2{
		Loan_id:       l.Loan_id,
		Checkout_date: l.Checkout_date,
		Due_date:      l.Due_date,
		Return_date:   l.Return_date,
		Status:        l.Status,
		Book_id:       b,
		User_id:       user,
		Admin_id:      admin,
		Penalties:     p,
		Created_at:    l.Created_at,
		Updated_at:    l.Updated_at,
	}
}

func ToDetailBookLoanResponseComplete(bl domain.BookLoan, b webresponse.BookResponseComplete, u domain.User, p domain.Penalties) webresponse.ListBookLoanResponse {
	book := webresponse.Book{
		Book_id: b.Book_id,
		Title:   b.Title,
		Foto:    b.Foto,
	}
	user := webresponse.User{
		User_id: u.User_id,
		Name:    u.Name,
		Foto:    u.Foto,
	}
	penalty := webresponse.Penalty{
		Penalty_id:     p.Penalty_id,
		Penalty_amount: p.Penalty_amount,
		Payment_status: p.Payment_status,
		Due_date:       p.Due_date.String(),
		Reason:         p.Reason,
	}
	return webresponse.ListBookLoanResponse{
		Loan_id:       bl.Loan_id,
		Checkout_date: bl.Checkout_date,
		Due_date:      bl.Due_date,
		Return_date:   bl.Return_date,
		Status:        bl.Status,
		Book:          book,
		User:          user,
		Admin_id:      bl.Admin_id,
		Penalties:     penalty,
		Created_at:    bl.Created_at,
		Updated_at:    bl.Updated_at,
	}
}

func ToPenaltiesResponse(p domain.Penalties) webresponse.PenaltiesResponse {
	return webresponse.PenaltiesResponse{
		Penalty_id:     p.Penalty_id,
		Loan_id:        p.Loan_id,
		Penalty_amount: p.Penalty_amount,
		Payment_status: p.Payment_status,
		Due_date:       p.Due_date,
		Reason:         p.Reason,
		Admin_id:       p.Admin_id,
		Created_at:     p.Created_at,
		Updated_at:     p.Updated_at,
	}
}
