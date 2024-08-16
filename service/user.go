package service

import (
	"context"
	"errors"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/repository"
	"github.com/tizor98/genpass/utils"
	"log"
)

func SaveNewPassword(password, forEntity string, user *entity.User) {
	pr := repository.PasswordRepository(context.Background())

	_, err := pr.Create(password, forEntity, user.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUsers() map[string]bool {
	ur := repository.UserRepository(context.Background())
	return ur.ListUsersNames()
}

func SetActive(username string) {
	ur := repository.UserRepository(context.Background())
	ur.SetActive(username)
}

func NewUser(username, password string) entity.User {
	password = utils.Encrypt(password)

	ur := repository.UserRepository(context.Background())

	userId, err := ur.Create(&entity.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	return ur.GetUser(userId)
}

func RemoveUser(username, pass string) error {
	ur := repository.UserRepository(context.Background())
	user := ur.GetUserByUsername(username)
	if user.Id == 0 {
		return errors.New("user not found")
	}

	if ok := utils.Compare(user.Password, pass); !ok {
		return errors.New("wrong password")
	}

	return ur.Delete(user.Username)
}
