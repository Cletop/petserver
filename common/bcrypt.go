package common

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password string) string {
	byte_password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(byte_password)
}

func VerifyPassword(originalPassword, newPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(newPassword))
}
