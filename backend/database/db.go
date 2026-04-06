package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
    var err error

    DB, err = sql.Open("sqlite", "./todo.db")
    if err != nil {
        panic(err)
    }

    createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT 0
	)
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}