package common

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), nil
}

// Equal 判断密码是否相同
func Equal(passEncrypt, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(passEncrypt), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
