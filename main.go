package main

import (
	"fmt"
	"net/http"

	"github.com/be/perpustakaan/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
}

func main(){
	fmt.Println("server running")
	server := InitializedServer()
	fmt.Println("jalan2")

	err := server.ListenAndServe()
	fmt.Println("jalan3")
	fmt.Println(err)

	helper.PanicIfError(err)
}