//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/be/perpustakaan/app"
	"github.com/be/perpustakaan/controller"
	"github.com/be/perpustakaan/repository"
	"github.com/be/perpustakaan/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

func InitializedServer() *http.Server {
	wire.Build(
		app.Cloudinary,
		app.NewDB,
		app.NewValidate,
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
		app.NewRouter,
		NewServer,
	)
	return nil
}