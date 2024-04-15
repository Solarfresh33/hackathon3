package models

import (
	"database/sql"
	// "fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenDateBase() {
	DB, _ = sql.Open("sqlite3", "DataBase.db")

}
func CreateDBUsers() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uid TEXT NOT NULL UNIQUE,
		adminusr TEXT NOT NULL,
		adminpswd TEXT NOT NULL,
		Admin INTEGER
	)`)

	if err != nil {
		print(err.Error())
	}
}

func CreateDBCommand() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS Command (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		IdColis TEXT,
		CodePostal INTEGER,
		Email TEXT,
		Name TEXT,
		Adresse VARCHAR(100),
		State TEXT,
		Date TEXT,
		EstimateTime TEXT,
		ville TEXT,
		Livre TEXT,
		Probleme TEXT,
		PointRelais TEXT
	)`)

	if err != nil {
		print(err.Error())
	}
}
