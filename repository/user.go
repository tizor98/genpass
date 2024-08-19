package repository

import (
	"context"
	"database/sql"
	"errors"
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
	SetActive(username string)
	SetNonActive(username string)
	ListUsersNames() map[string]bool
	Delete(username string) error
}

func UserRepository(ctx context.Context) UserRepo {
	return &userRepo{db: getDb(ctx)}
}

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) Delete(username string) error {
	result, err := u.db.Exec("DELETE FROM users WHERE username = ?", username)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if rows != 1 {
		return errors.New("user not found")
	}

	return nil
}

func (u *userRepo) ListUsersNames() map[string]bool {
	rows, err := u.db.Query("SELECT username, is_active FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(rows, "ListUsersNames")

	output := make(map[string]bool)
	for rows.Next() {
		var userName string
		var isActive bool
		if err = rows.Scan(&userName, &isActive); err != nil {
			log.Fatal(err)
		}
		output[userName] = isActive
	}

	return output
}

func (u *userRepo) GetUser(id int64) entity.User {
	row, err := u.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(row, "GetUser")

	var foundUser entity.User
	if row.Next() {
		if err := row.Scan(&foundUser.Id, &foundUser.Username, &foundUser.Password, &foundUser.IsActive, &foundUser.CreatedAt, &foundUser.UpdatedAt); err != nil {
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
        INSERT INTO users (username, password) VALUES (?, ?)
        `, user.Username, user.Password)

	if err != nil {
		log.Fatal(err)
	}
	return result.LastInsertId()
}

func (u *userRepo) GetActive() entity.User {
	row, err := u.db.Query("SELECT * FROM users WHERE is_active = true")
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(row, "GetUserByUsername")

	var user entity.User
	scanOneStruct(row, &user)

	return user
}

func (u *userRepo) SetActive(username string) {
	_, err := u.db.Exec("UPDATE users SET is_active = false WHERE is_active = true")
	if err != nil {
		log.Fatal(err)
	}
	_, err = u.db.Exec("UPDATE users SET is_active = true WHERE username = ?", username)
	if err != nil {
		log.Fatal(err)
	}
}

func (u *userRepo) SetNonActive(username string) {
	_, err := u.db.Exec("UPDATE users SET is_active = false WHERE username = ?", username)
	if err != nil {
		log.Fatal(err)
	}
}
