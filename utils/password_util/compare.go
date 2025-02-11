package password_util

import (
	"golang.org/x/crypto/bcrypt"
)

func Compare(hashedPassword string, rawPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	if err != nil {
		return false, err
	}

	return true, nil
}
