package app

import (
	"fmt"
	"os"

	"github.com/be/perpustakaan/helper"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

func Cloudinary() *cloudinary.Cloudinary{
	errEnv := godotenv.Load()
	helper.PanicIfError(errEnv)

	fmt.Println(os.Getenv("CLOUDINARY_URL"))
	cld, errCld := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	helper.PanicIfError(errCld)
	return cld
}