package helper

import (
	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webresponse"
)

func ToUserResponse(user domain.User) webresponse.UserResponse {
	return webresponse.UserResponse{
	Name: user.Name,
	Email: user.Email,
	Level: user.Level,
	Is_enabled: user.Is_enabled,
	Gender: user.Gender,
	Telp : user.Telp,
	Birthdate :user.Birthdate,
	Address : user.Address,
	Foto : user.Foto,
	Batas :user.Batas,
	}
}

func ToUserResponses(users []domain.User) []webresponse.UserResponse{
	var userResponses []webresponse.UserResponse

	for _,user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}