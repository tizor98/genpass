package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encrypt(target string) string {
	encryptData, err := bcrypt.GenerateFromPassword([]byte(target), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(encryptData)
}

func Compare(encryptedStr, plaintStr string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedStr), []byte(plaintStr)); err != nil {
		return false
	}
	return true
}

func EncryptWithKeys(target string, keys ...string) string {
	var sb strings.Builder

	for _, key := range keys {
		sb.WriteString(key)
	}

	if sb.Len() < 32 {
		sliceOfZeros := make([]byte, 32-sb.Len())
		sb.Write(sliceOfZeros)
	}

	block, err := aes.NewCipher([]byte(sb.String()[:32]))
	if err != nil {
		log.Fatal(err)
	}

	output := make([]byte, len(target))

	cfb := cipher.NewCFBEncrypter(block, bytes)
	cfb.XORKeyStream(output, []byte(target))
	return base64.StdEncoding.EncodeToString(output)
}

func DecryptWithKeys(target string, keys ...string) string {
	var sb strings.Builder

	for _, key := range keys {
		sb.WriteString(key)
	}

	if sb.Len() < 32 {
		sliceOfZeros := make([]byte, 32-sb.Len())
		sb.Write(sliceOfZeros)
	}

	input, err := base64.StdEncoding.DecodeString(target)
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher([]byte(sb.String()[:32]))
	if err != nil {
		log.Fatal(err)
	}

	output := make([]byte, len(input))

	cfb := cipher.NewCFBDecrypter(block, bytes)
	cfb.XORKeyStream(output, input)
	return string(output)
}
