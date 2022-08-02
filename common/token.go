package common

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user_id uint, username string) (string, error) {
	var err error

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"authorized": true,
		"user_id":    user_id,
		"username":   username,
		"exp":        time.Now().Add(time.Minute * 15).Unix(),
	})

	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
