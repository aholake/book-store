package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "./bookstore.db"

var DB *sql.DB

func Connect() error {
	var err error
	DB, err = sql.Open("sqlite3", file)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}
	var version string

	if err = DB.QueryRow("SELECT SQLITE_VERSION()").Scan(&version); err != nil {
		return err
	}

	fmt.Println("Connected to database, version", version)
	return nil
}
