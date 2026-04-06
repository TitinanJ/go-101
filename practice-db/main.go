package main

import (
	"database/sql"
	"fmt"
	"with-db/dbo/Person"

	_ "modernc.org/sqlite"
)

func main() {
    db, err := sql.Open("sqlite", "./database.sqlite3")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    if _, err := db.Exec(`DROP TABLE IF EXISTS users;
        CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    )`); err != nil {
        panic(err)
    }

    if _, err := db.Exec(`INSERT INTO users (name)
        VALUES (?), (?)`, "Alice", "Bob"); err != nil {
        panic(err)
    }

    rows, err := db.Query(`SELECT id, name FROM users`)
    if err != nil {
        panic(err)
    }

    for rows.Next() {
        person := new(person.Person)
        if err = rows.Scan(&person.ID, &person.Name); err != nil {
            panic(err)
        }

        fmt.Println(person.ID, person.Name)
    }
}