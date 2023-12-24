package helper

// func GenerateJWT(user webrequest.UserGenereteToken) (string, error) {
// 	claims := jwt.MapClaims{
// 		"id" : user.Id,
// 		"email": user.Email,
// 		"exp":  time.Now().Add(time.Hour * 24).Unix(),

// 	}
// 	secret:=[]byte(os.Getenv("SECRET_KEY"))
// 	fmt.Println("jalan")
// 	fmt.Println(secret)
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

// 	tokenString,err := token.SignedString(secret)

// 	if err != nil {
// 		fmt.Println("err")
// 		fmt.Println(err)
// 		return "", err
// 	}
// 	fmt.Println(tokenString)
// 	return tokenString, nil
// }