package models

import (
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
	Email        string
	Name         string
	Connected 	bool
}

func CreatePackage(adresse string, IdColis string, Email string, Name string, codepostal int, Date string, State string, Estimatetime string, ville string, pointrelais string) (UsersPackage, error) {
	var createpackage UsersPackage
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var uid string
	for i := 0; i < 16; i++ {
		uid += string(charset[seededRand.Intn(len(charset))])
	}
	println(uid)
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
	currentTime := time.Now().In(parisLocation)
	rand.Seed(time.Now().UnixNano())
	randomNumber := time.Duration(rand.Intn(9) + 4)
	futureTime := currentTime.Add(randomNumber * 24 * time.Hour)
	formattedFutureTime := futureTime.Format("Monday 02 January")
	_, err = DB.Exec("INSERT INTO command (adresse, codepostal, IdColis, Date, State, Estimatetime, ville, PointRelais, Probleme, Livre, Email, Name) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", adresse, codepostal, uid, currentTime.Format("Monday 02 January"), "En préparation", formattedFutureTime, ville, pointrelais, "Non", "Non", Email, Name)
	if err != nil {
		panic(err)
	}
	createpackage.Adresse = adresse
	createpackage.CodePostal = codepostal
	createpackage.Ville = ville
	createpackage.PointRelais = pointrelais
	createpackage.Email = Email
	createpackage.Name = Name
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

	rows, err := DB.Query("SELECT date FROM command WHERE idcolis=?", uid)

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
	Connected bool
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
	_, err := DB.Exec("INSERT INTO Users (adminusr, adminpswd, uid) VALUES (?, ?, ?)", "admin", "admin", UserAdmin.Uid)
	if err != nil {
		panic(err)
	}
}


func GetEmailandNAme(idcolis string) UsersPackage{
	rows, err := DB.Query("SELECT Name, Email FROM command WHERE IdColis = ?", idcolis)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var GetMail UsersPackage
	for rows.Next() {
		err := rows.Scan(&GetMail.Name, &GetMail.Email)
		if err != nil {
			panic(err)
		}
	}
	println("Email : " + GetMail.Email)
	println("Name : " + GetMail.Name)
	return GetMail
}

func ExistAccount(Pseudo string) (bool, string, string) {
	rows, _ := DB.Query("SELECT adminusr , adminpswd, uid FROM Users")
	defer rows.Close()
	for rows.Next() {
		var each_pseudo string
		var each_psswd string
		var uid string
		_ = rows.Scan(&each_pseudo, &each_psswd, &uid)
		if each_pseudo == Pseudo {
			return true, each_psswd, uid
		}
	}
	return false, "", "oui"
}