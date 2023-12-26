package konversi

import (
	"strconv"

	"github.com/be/perpustakaan/exception"
)

func StrToInt(str string, name string) int {
	conv, err := strconv.Atoi(str)
	if err != nil {
		errr := name + " must greater than zero"
		panic(exception.CustomEror{Code: 400, Error: errr})
	}
	if conv <= 0 {
		errr := name + " must greater than zero"
		panic(exception.CustomEror{Code: 400, Error: errr})
	}
	return conv
}
