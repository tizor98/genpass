package repository

import (
	"context"
	"database/sql"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/utils"
	"log"
)

type PasswordRepo interface {
	GetPassword(id int64) entity.Password
	GetPasswordByForAndUsername(forEntity, username string) entity.Password
	ForPasswordsListByUsername(username string) []string
	ExistsPasswordForEntity(forEntity, username string) bool
	Create(password, forEntity string, userId int64) (int64, error)
	DeleteByUsername(username string) error
}

func PasswordRepository(ctx context.Context) PasswordRepo {
	return &passwordRepo{db: getDb(ctx)}
}

type passwordRepo struct {
	db *sql.DB
}

func (p passwordRepo) DeleteByUsername(username string) error {
	_, err := p.db.Exec("DELETE FROM passwords WHERE user_id = (SELECT id FROM users WHERE username = ?)", username)
	return err
}

func (p passwordRepo) GetPassword(id int64) entity.Password {
	row, err := p.db.Query("SELECT * FROM passwords WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(row, "GetPassword")

	var pass entity.Password
	scanOneStruct(row, &pass)

	return pass
}

func (p passwordRepo) GetPasswordByForAndUsername(forEntity, username string) entity.Password {
	row, err := p.db.Query(`
		SELECT p.* 
		FROM passwords p 
		    JOIN main.users u ON u.id = p.user_id 
		WHERE u.username = ? AND p.for_entity = ?
		`, username, forEntity)
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(row, "GetPasswordByForAndUsername")

	var pass entity.Password
	scanOneStruct(row, &pass)

	return pass
}

func (p passwordRepo) ForPasswordsListByUsername(username string) []string {
	row, err := p.db.Query(`
		SELECT for_entity 
		FROM passwords p 
		    JOIN main.users u ON u.id = p.user_id 
		WHERE u.username = ?
		`, username)
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(row, "ForPasswordsListByUsername")

	forEntities := make([]string, 0)
	for row.Next() {
		var forEntity string
		if err := row.Scan(&forEntity); err != nil {
			log.Fatal(err)
		}
		forEntities = append(forEntities, forEntity)
	}
	return forEntities
}

func (p passwordRepo) ExistsPasswordForEntity(forEntity, username string) bool {
	row, err := p.db.Query(`
		SELECT COUNT(*) 
		FROM passwords p 
		    JOIN main.users u ON u.id = p.user_id 
		WHERE u.username = ? AND p.for_entity = ?
		`, username, forEntity)
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Close(row, "ForPasswordsListByUsername")

	count := -1
	if row.Next() {
		if err := row.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	return count > 0
}

func (p passwordRepo) Create(password, forEntity string, userId int64) (int64, error) {
	result, err := p.db.Exec(`
		INSERT INTO passwords (user_id, password, for_entity) VALUES (?, ?, ?);
		`, userId, password, forEntity)
	if err != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}
