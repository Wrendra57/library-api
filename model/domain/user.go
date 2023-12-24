package domain

import (
	"time"
)

type User struct {
	User_id int
	Name string
	Email string
	Password string
	Level string
	Is_enabled bool
	Gender string
	Telp string
	Birthdate time.Time
	Address string
	Foto string
	Batas int
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}