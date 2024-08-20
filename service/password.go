package service

import (
	"context"
	"errors"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/repository"
	"github.com/tizor98/genpass/utils"
	"log"
	"math/rand/v2"
	"slices"
	"strings"
)

type PassType byte

const (
	PassTypeAll = PassType(iota)
	PassTypeCapitalAndLower
	PassTypeCapitalAndLowerAndNumber
)

var (
	CapitalCase       = []byte{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90}
	LowerCase         = []byte{97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122}
	NumberCase        = []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	SpecialCharacters = []byte{33, 64, 35, 36, 37, 94, 38, 42, 40, 41, 95, 43, 45, 61, 123, 124, 125, 91, 93, 59, 58, 44, 46, 47, 63, 126}
	PasswordLength    = 20
)

var (
	ErrPasswordNotFound = errors.New("password not found")
)

func GetPassword(forEntity, username, userPass string) (string, error) {
	pr := repository.PasswordRepository(context.Background())
	encryptedForEntity := utils.EncryptWithKey(forEntity, []string{userPass})

	pass := pr.GetPasswordByForAndUsername(encryptedForEntity, username)

	if pass.Id == 0 || pass.Password == "" {
		return "", ErrPasswordNotFound
	}

	password := utils.DecryptWithKey(pass.Password, []string{encryptedForEntity, username})

	return password, nil
}

func SaveNewPassword(pass, forEntity string, user *entity.User, userPass string) {
	pr := repository.PasswordRepository(context.Background())

	encryptedForEntity := utils.EncryptWithKey(forEntity, []string{userPass})
	password := utils.EncryptWithKey(pass, []string{encryptedForEntity, user.Username})

	_, err := pr.Create(password, encryptedForEntity, user.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func NewPassword(ctx context.Context) string {
	length := ctx.Value(utils.NewFlagPassLength)
	if length != nil {
		PasswordLength = length.(int)
	}

	mode := ctx.Value(utils.NewFlagPassType).(PassType)
	source := getSource(mode)

	return generatePassword(PasswordLength, source)
}

func generatePassword(length int, source []byte) string {
	var sb strings.Builder

	for sb.Len() < length {
		rand.Shuffle(len(source), func(i, j int) {
			source[i], source[j] = source[j], source[i]
		})

		index := len(source)
		if length-sb.Len() < len(source) {
			index = length - sb.Len()
		}
		sb.Write(source[:index])
	}

	return sb.String()
}

func getSource(mode PassType) []byte {
	var source []byte
	switch mode {
	case PassTypeAll:
		source = append(append(append(CapitalCase, LowerCase...), NumberCase...), SpecialCharacters...)
		break
	case PassTypeCapitalAndLowerAndNumber:
		source = append(append(CapitalCase, LowerCase...), NumberCase...)
		break
	case PassTypeCapitalAndLower:
		source = append(CapitalCase, LowerCase...)
		break
	}
	return source
}

func GetAllPasswords(username string, userPass string) []string {
	pr := repository.PasswordRepository(context.Background())
	forEntitySlice := pr.ForPasswordsListByUsername(username)

	for i, forEntity := range forEntitySlice {
		forEntitySlice[i] = utils.DecryptWithKey(forEntity, []string{userPass})
	}

	slices.Sort(forEntitySlice)
	return forEntitySlice
}
