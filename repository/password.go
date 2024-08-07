package repository

import (
	"context"
	"database/sql"
	"github.com/tizor98/genpass/entity"
)

type PasswordRepo interface {
	GetPassword(id uint64) entity.Password
	GetPasswordByForAndUsername(forEntity, username string) entity.Password
	ForPasswordsListByUsername(username string) []string
	ExistsPasswordForEntity(forEntity string) bool
	Create(user entity.Password) error
}

func PasswordRepository(ctx context.Context) UserRepo {
	return &passwordRepo{db: getDb(ctx)}
}

type passwordRepo struct {
	db *sql.DB
}

func (p passwordRepo) GetUser(id uint64) entity.User {
	//TODO implement me
	panic("implement me")
}

func (p passwordRepo) GetUserByUsername(username string) entity.User {
	//TODO implement me
	panic("implement me")
}

func (p passwordRepo) ExistByUsername(username string) bool {
	//TODO implement me
	panic("implement me")
}

func (p passwordRepo) Create(user entity.User) (id uint64, err error) {
	//TODO implement me
	panic("implement me")
}
