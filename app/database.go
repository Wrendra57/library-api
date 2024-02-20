package app

import (
	"database/sql"
	"os"

	"time"

	"github.com/be/perpustakaan/helper"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	errEnv := godotenv.Load()
	helper.PanicIfError(errEnv)
	db_url := os.Getenv("DATABASE_URL") + "?parseTime=true"

	db, err := sql.Open("mysql", db_url)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
