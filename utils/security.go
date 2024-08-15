package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Encrypt(target string) string {
	encryptData, err := bcrypt.GenerateFromPassword([]byte(target), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(encryptData)
}

func Compare(target, source string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(target), []byte(source)); err != nil {
		return false
	}
	return true
}
