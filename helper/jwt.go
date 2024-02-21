package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/be/perpustakaan/model/webrequest"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user webrequest.UserGenereteToken) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"level": user.Level,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	// fmt.Println("expired ==>", time.Now().Add(time.Hour*24).Unix())
	secret := []byte(os.Getenv("SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}

func ParseJWT(tokenString string) (*webrequest.UserGenereteToken, error) {
	// Parse the token
	secret := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(tokenString, &webrequest.UserGenereteToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	fmt.Println("parse    ", token)
	if err != nil {
		// fmt.Println("Error parsing with claims:")
		return &webrequest.UserGenereteToken{}, err
	}

	// Verify the token
	if !token.Valid {
		// fmt.Println("token is not valid")
		return &webrequest.UserGenereteToken{}, fmt.Errorf("invalid token")
	}

	// Extract custom claims
	claims, ok := token.Claims.(*webrequest.UserGenereteToken)
	if !ok {
		// fmt.Println("err claims")
		return &webrequest.UserGenereteToken{}, fmt.Errorf("failed to parse custom claims")
	}
	// fmt.Println(claims)
	return claims, nil
}
