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

// func DeleteDB() {
// 	_, err := DB.Exec("DELETE FROM u") //le nom de la table a delete
// 	if err != nil {
// 		fmt.Println("Erreur lors de l'exécution de la requête DELETE:", err)
// 		return
// 	}

// 	fmt.Println("Table vidée avec succès.")
// }


func CreateDBCommand() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS Command (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		IdColis INTEGER,
		CodePostal INTEGER,
		Adresse VARCHAR(100),
		State TEXT,
		Date TEXT,
		EstimateTime TIMESTAMP
	)`)

	if err != nil {
		print(err.Error())
	}
}

