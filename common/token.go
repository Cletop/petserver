package common

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(user_id uint, username string) (string, error) {
	var err error

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"user_id":  user_id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"jti":      user_id,
		"sub":      "petserver_auth",
	})

	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(token string) (uint, string, error) {
	var err error
	var user_id uint
	var username string

	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims := at.Claims.(jwt.MapClaims)
	user_id = uint(claims["user_id"].(float64))
	username = claims["username"].(string)

	return user_id, username, nil
}
