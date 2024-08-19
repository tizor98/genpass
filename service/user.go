package service

import (
	"context"
	"errors"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/repository"
	"github.com/tizor98/genpass/utils"
	"log"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrWrongPassword = errors.New("wrong password")
	ErrUserExists    = errors.New("user already exists")
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

func SetActive(username, pass string) error {
	ur := repository.UserRepository(context.Background())
	user := ur.GetUserByUsername(username)
	if user.Id == 0 {
		return ErrUserNotFound
	}

	if ok := utils.Compare(user.Password, pass); !ok {
		return ErrWrongPassword
	}
	ur.SetActive(username)

	return nil
}

func NewUser(username, password string) (entity.User, error) {
	ur := repository.UserRepository(context.Background())
	oldUser := ur.GetUserByUsername(username)
	if oldUser.Id > 0 {
		return entity.User{}, ErrUserExists
	}

	password = utils.Encrypt(password)

	userId, err := ur.Create(&entity.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	return ur.GetUser(userId), nil
}

func RemoveUser(username, pass string) error {
	ur := repository.UserRepository(context.Background())
	user := ur.GetUserByUsername(username)
	if user.Id == 0 {
		return ErrUserNotFound
	}

	if ok := utils.Compare(user.Password, pass); !ok {
		return ErrWrongPassword
	}

	ps := repository.PasswordRepository(context.Background())
	err := ps.DeleteByUsername(username)
	if err != nil {
		return err
	}

	return ur.Delete(user.Username)
}

func GetActive() entity.User {
	ur := repository.UserRepository(context.Background())
	return ur.GetActive()
}

func SetNonActive(username, pass string) error {
	ur := repository.UserRepository(context.Background())
	user := ur.GetUserByUsername(username)
	if user.Id == 0 {
		return ErrUserNotFound
	}

	if ok := utils.Compare(user.Password, pass); !ok {
		return ErrWrongPassword
	}
	ur.SetNonActive(username)
	return nil
}
