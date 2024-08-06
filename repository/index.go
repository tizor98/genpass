package repository

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path"
)

const (
	dbDir  = ".genpass"
	dbName = "general.db"
)

var db *sql.DB

func init() {
	if db == nil {
		getDb(context.Background())
	}
	initDatabase()
}

func getDb(ctx context.Context) *sql.DB {
	if db != nil {
		err := db.PingContext(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return db
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(path.Join(dir, dbDir), 0755)
	if err != nil {
		log.Fatal(err)
	}

	dbPath := path.Join(dir, dbDir, dbName)

	dataBase, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	db = dataBase
	return db
}

func initDatabase() {
	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS main.users (
            id INTEGER PRIMARY KEY, 
            username VARCHAR(20) NOT NULL UNIQUE,
            name VARCHAR(64) NOT NULL,
            surname VARCHAR(64),
            created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        );
        
        CREATE TRIGGER IF NOT EXISTS update_users_updated_at
        AFTER UPDATE ON users 
        WHEN old.updated_at <> CURRENT_TIMESTAMP
        BEGIN
            UPDATE users SET updated_at = CURRENT_TIMESTAMP
            WHERE id = old.id;
        END;
end;
    `); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS passwords (
            id INTEGER PRIMARY KEY,
            target VARCHAR NOT NULL,
            password VARCHAR NOT NULL,
            userId INTEGER NOT NULL,
            created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY(userId) REFERENCES users(id)
        );

        CREATE TRIGGER IF NOT EXISTS update_passwords_updated_at
        AFTER UPDATE ON passwords
        WHEN old.updated_at <> CURRENT_TIMESTAMP
        BEGIN
            UPDATE passwords SET updated_at = CURRENT_TIMESTAMP
            WHERE id = old.id;
        END;
    `); err != nil {
		log.Fatal(err)
	}
}
