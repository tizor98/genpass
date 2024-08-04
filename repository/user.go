package repository

import (
    "context"
    "database/sql"
    "os/user"
)

type UserRepo interface {
    GetUser(id string) user.User
    GetUserByUsername(username string) user.User
    ExistByUsername(username string) bool
    Create(user user.User) error
}

func UserRepository(ctx context.Context) UserRepo {
    return &userRepo{db: getDb()}
}

type userRepo struct {
    db *sql.DB
}

func (u *userRepo) GetUser(id string) user.User {
    //TODO implement me
    panic("implement me")
}

func (u *userRepo) GetUserByUsername(username string) user.User {
    //TODO implement me
    panic("implement me")
}

func (u *userRepo) ExistByUsername(username string) bool {
    //TODO implement me
    panic("implement me")
}

func (u *userRepo) Create(user user.User) error {
    //TODO implement me
    panic("implement me")
}
