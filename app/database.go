package app

import (
	"database/sql"

	"time"

	"github.com/be/perpustakaan/helper"
)

func NewDB() *sql.DB{
	db,err:=sql.Open("mysql","root:password@tcp(localhost:3306)/library?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}