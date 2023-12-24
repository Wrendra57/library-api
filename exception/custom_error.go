package exception

type CustomEror struct {
	Code  int
	Error string
}

func NewCustomEror(code int,error string) CustomEror {
	return CustomEror{
		Code: code,
		Error: error}
}
