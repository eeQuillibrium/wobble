package utils

import "golang.org/x/crypto/bcrypt"

func HashPass(pass string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
