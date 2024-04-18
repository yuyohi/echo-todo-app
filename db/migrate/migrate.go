package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../todo_app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlFile, err := os.ReadFile("setupDB.sql")
	if err != nil {
		log.Fatal(err)
	}
	sqlString := string(sqlFile)

	_, err = db.Exec(sqlString)
	if err != nil {
		log.Fatal(err)
	}
}
