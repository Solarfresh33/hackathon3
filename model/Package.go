package models

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

type UsersPackage struct {
	Id           int
	Adresse      string
	CodePostal   int
	IdColis      string
	State        string
	Date         string
	EstimateTime string
	Ville        string
	PointRelais  string
}

func CreatePackage(adresse string, IdColis string, codepostal int, Date string, State string, Estimatetime string, ville string, pointrelais string) (UsersPackage, error) {
	var createpackage UsersPackage
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var uid string
	for i := 0; i < 16; i++ {
		uid += string(charset[seededRand.Intn(len(charset))])
	}
	println(uid)
	h := md5.New()
	idStr := uid
	h.Write([]byte(idStr))
	idStr = hex.EncodeToString(h.Sum(nil))
	parisLocation, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		panic(err)
	}

	rows, err := DB.Query("SELECT date FROM command", Date)
	if err != nil {
		panic(err)
	}
	var DateEstimate string
	for rows.Next() {
		err := rows.Scan(&DateEstimate)
		if err != nil {
			panic(err)
		}
	}
	println(DateEstimate)
	currentTime := time.Now().In(parisLocation)
	futureTime := currentTime.Add(4 * 24 * time.Hour)
	formattedFutureTime := futureTime.Format("Monday 02 January")
	println(formattedFutureTime)
	println("quoicoucrampté")
	_, err = DB.Exec("INSERT INTO command (adresse, codepostal, IdColis, Date, State, Estimatetime, ville, PointRelais, Probleme, Livre) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", adresse, codepostal, idStr, currentTime.Format("Monday 02 January"), "En préparation", formattedFutureTime, ville, pointrelais, "Non", "Non")
	if err != nil {
		panic(err)
	}
	createpackage.Adresse = adresse
	createpackage.CodePostal = codepostal
	createpackage.Ville = ville
	createpackage.PointRelais = pointrelais

	dates, err := GetDate(uid)
	if err != nil {
		return createpackage, err
	}

	if len(dates) > 0 {
		createpackage.Date = dates[0].Dates
	}

	return createpackage, nil
}

type Date struct {
	Dates string
	Uid   string
}

func GetDate(uid string) ([]Date, error) {
	var dates []Date

	rows, err := DB.Query("SELECT date FROM command", uid)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var date Date
	for rows.Next() {
		err := rows.Scan(&date.Dates)
		if err != nil {
			panic(err)
		}
		dates = append(dates, date)
	}
	return dates, nil
}

func UpdatePackageState(idColis string, newState string, newpointRelais string, livre string, probleme string) error {
	_, err := DB.Exec("UPDATE command SET State = ?, PointRelais = ?, Livre = ?, Probleme = ? WHERE idColis = ?", newState, newpointRelais, livre, probleme, idColis)
	if err != nil {
		panic(err)
	}
	return nil
}

type Users struct {
	Id        int
	Uid       string
	Adminusr  string
	Adminpswd string
	Admin     int
}

func CreateUserAdmin(username string, password string, UID string, admin int) (Users, error) {
	var UserAdmin Users
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var uid string
	for i := 0; i < 16; i++ {
		uid += string(charset[seededRand.Intn(len(charset))])
	}
	_, err := DB.Exec("INSERT INTO Users(adminusr, adminpswd, Admin, uid) VALUES(?, ?, 1, ?)", username, password, uid)
	if err != nil {
		panic(err)
	}
	UserAdmin.Adminusr = username
	UserAdmin.Adminpswd = password
	UserAdmin.Uid = uid
	println("user crée avec come uid ", UserAdmin.Uid)
	return UserAdmin, nil
}

func FirstUserAdmin() {
	var UserAdmin Users
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		UserAdmin.Uid += string(charset[seededRand.Intn(len(charset))])
	}
	_, err := DB.Exec("INSERT INTO Users (adminusr, adminpswd, Admin, uid) VALUES (?, ?, 1, ?)", "admin", "admin", UserAdmin.Uid)
	if err != nil {
		panic(err)
	}
}
