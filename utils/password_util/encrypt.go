package password_util

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(rawPassword string) string {
	encryptedBytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(encryptedBytes)
}
