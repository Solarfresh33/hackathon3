package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

type UsersPackage struct {
	Id         int
	Adresse    string
	CodePostal int
	IdUser     string
	State      string
	Date       string
	// EstimateTime *timestamp.Timestamp
}

func GetPackageByID(id string, CodePostal int) (UsersPackage, error) {
	var colis UsersPackage
	err := DB.QueryRow("SELECT id, Adresse, CodePostal, IdUser, State, Date FROM command WHERE id=?", id).
		Scan(&colis.Id, &colis.Adresse, &colis.CodePostal, &colis.IdUser, &colis.State, &colis.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return colis, fmt.Errorf("no package found with ID %s", id)
		}
		return colis, err
	}

	return colis, nil
}

func CreatePackage(adresse string, IdUser string, codepostal int) (UsersPackage, error) {
	var createpackage UsersPackage
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var uid string
	for i := 0; i < 16; i++ {
		uid += string(charset[seededRand.Intn(len(charset))])
	}
    h := md5.New()
	idStr := uid
	h.Write([]byte(idStr))
	idStr = hex.EncodeToString(h.Sum(nil))
    println("2",codepostal,adresse)
    _, err := DB.Exec("INSERT INTO command (adresse, codepostal, IdColis) VALUES (?, ?, ?)", adresse, codepostal, idStr)
	createpackage.Adresse = adresse
	createpackage.CodePostal = codepostal
	if err != nil {
		return createpackage, err
	}
	return createpackage, nil
}
