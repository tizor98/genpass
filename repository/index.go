package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const (
	dbDir  = "/genpass"
	dbName = "/general.db"
)

var db *sql.DB

func init() {
	if db == nil {
		getDb()
	}
	initDatabase()
}

func getDb() *sql.DB {
	if db != nil {
		return db
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(dir+dbDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	dbPath := dir + dbDir + dbName

	dataBase, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	db = dataBase
	defer db.Close()
	return db
}

func initDatabase() {
	sta, err := db.Prepare(`
        CREATE IF NOT EXIST TABLE users (
               id PRIMARY KEY, 
               username string UNIQUE KEY,
               name string,
               surname string,
               created_at timestamptz not null default now()
               updated_at timestamptz not null default now() on update now()
        );
    `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = sta.Exec()
	if err != nil {
		log.Fatal(err)
	}
}
