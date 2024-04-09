package models

import (
	"log"
	"time"
)

type Date struct {
	Dates string
	Uid   string
}

func CreateDate() error {
	parisLocation, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		return err
	}

	currentTime := time.Now().In(parisLocation).Format("2006-01-02 15:04:05")

	statement := "INSERT INTO Command (Date) VALUES (?)"

	_, err = DB.Exec(statement, currentTime)
	if err != nil {
		log.Println("Erreur date :", err)
		return err
	}
	return nil
}

func GetDate(uid string) ([]Date, error) {
	var dates []Date

	rows, err := DB.Query("SELECT Date FROM Command WHERE Command.UIDSalé = ?", uid)

	if err != nil {
		log.Println("Erreur lors de la récupération de la date :", err)
		return dates, err
	}
	defer rows.Close()

	for rows.Next() {

		var date Date
		err := rows.Scan(&date.Dates)
		if err != nil {

			log.Println("Erreur lors du scan de la date :", err)
			return dates, err
		}
		dates = append(dates, date)
	}
	return dates, nil
}
