package common

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// access token expire time
var AccessTokenExpireTime = time.Now().Add(time.Minute * 10)    // 10 minutes
var RefreshTokenExpireTime = time.Now().Add(time.Hour * 24 * 7) // 7 days

type Auth interface {
	CreateTokenSignature() (string, error)
	SetSecret(string)
}
type CustomClaims struct {
	UserID      uint     `json:"user_id"`
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

func (claims *CustomClaims) CreateTokenSignature() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
func (claims *CustomClaims) SetSecret(secret string) {
	os.Setenv("JWT_SECRET", secret)
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

func CreateAccessToken(user_id uint, username string) (string, error) {
	return NormalCreateToken(user_id, username, AccessTokenExpireTime, []string{"access"})
}
func CreateRefreshToken(user_id uint, username string) (string, error) {
	return NormalCreateToken(user_id, username, RefreshTokenExpireTime, []string{"refresh"})
}
func NormalCreateToken(user_id uint, username string, expiresAt time.Time, permissions []string) (string, error) {
	claims := &CustomClaims{
		UserID:      user_id,
		Username:    username,
		Permissions: append(permissions, "auth_center"),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "petserver",
			Subject:   "petserver auth",
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	return claims.CreateTokenSignature()
}

func CreateRenewableToken(user_id uint, username string) (string, string, error) {
	access_token, err := CreateAccessToken(user_id, username)
	if err != nil {
		return "", "", err
	}
	refresh_token, err := CreateRefreshToken(user_id, username)
	if err != nil {
		return "", "", err
	}
	return access_token, refresh_token, nil
}
