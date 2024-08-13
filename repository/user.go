package repository

import (
	"context"
	"database/sql"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/utils"
	"log"
)

type UserRepo interface {
	GetUser(id int64) entity.User
	GetUserByUsername(username string) entity.User
	ExistByUsername(username string) bool
	Create(user *entity.User) (id int64, err error)
	GetActive() entity.User
}

func UserRepository(ctx context.Context) UserRepo {
	return &userRepo{db: getDb(ctx)}
}

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) GetUser(id int64) entity.User {
	row, err := u.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(row, "GetUser")

	var foundUser entity.User
	if row.Next() {
		if err := row.Scan(&foundUser.Id, &foundUser.Name, &foundUser.Surname, &foundUser.Username, &foundUser.Password, &foundUser.CreatedAt, &foundUser.UpdatedAt); err != nil {
			log.Fatal(err)
		}
	}
	return foundUser
}

func (u *userRepo) GetUserByUsername(username string) entity.User {
	row, err := u.db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(row, "GetUserByUsername")

	var user entity.User
	scanOneStruct(row, &user)

	return user
}

func (u *userRepo) ExistByUsername(username string) bool {
	row, err := u.db.Query("SELECT COUNT(*) FROM users WHERE username = ?", username)
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(row, "ExistByUsername")

	count := -1
	if row.Next() {
		err := row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		return count == 1
	}
	return false
}

func (u *userRepo) Create(user *entity.User) (id int64, err error) {
	result, err := u.db.Exec(`
        INSERT INTO users (username, name, surname, password) VALUES (?, ?, ?, ?)
        `, user.Username, user.Name, user.Surname, user.Password)

	if err != nil {
		log.Fatal(err)
	}
	return result.LastInsertId()
}

func (u *userRepo) GetActive() entity.User {
	row, err := u.db.Query("SELECT * FROM users WHERE active = true")
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(row, "GetUserByUsername")

	var user entity.User
	scanOneStruct(row, &user)

	return user
}
