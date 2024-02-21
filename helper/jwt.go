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

	secret := []byte(os.Getenv("SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (*webrequest.UserGenereteToken, error) {
	// Parse the token
	secret := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(tokenString, &webrequest.UserGenereteToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {

		return &webrequest.UserGenereteToken{}, err
	}

	// Verify the token
	if !token.Valid {

		return &webrequest.UserGenereteToken{}, fmt.Errorf("invalid token")
	}

	// Extract custom claims
	claims, ok := token.Claims.(*webrequest.UserGenereteToken)
	if !ok {

		return &webrequest.UserGenereteToken{}, fmt.Errorf("failed to parse custom claims")
	}

	return claims, nil
}
