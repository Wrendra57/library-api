package webresponse

import "time"

type UserResponse struct {
	User_id int			`json:"user_id"`
	Name string		`json:"name"`
	Email string	`json:"email"`
	
	Level string 	`json:"level"`
	Is_enabled bool `json:"is_enabled"`
	Gender string `json:"gender"`
	Telp string `json:"telp"`
	Birthdate time.Time `json:"birthdate"`
	Address string `json:"address"`
	Foto string `json:"foto"`
	Batas int `json:"batas"`

}

type LoginResponse struct {
	Token string `json:"token"`
}