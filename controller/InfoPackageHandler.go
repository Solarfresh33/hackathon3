package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
	"strings"
)

type PackageInfo struct {
	CodePostal   string
	Adresse      string
	State        string
	Date         string
	EstimateTime string
	Ville        string
	Idcolis      string
	PointRelais  string
	Probleme     string
	Livre        string
	Connected bool
}

func InfoPackageHandler(w http.ResponseWriter, r *http.Request) {
	var packageInfo PackageInfo
	cookie, _ := r.Cookie("User")
	if cookie == nil {
		packageInfo.Connected = false
		println("Connected = false")
	} else {
		packageInfo.Connected = true
		println("Connected = true")
	
	}
	GetURLID := strings.Split(r.URL.Path, "/")[2]

	packageInfo.Idcolis = GetURLID
	if GetURLID != "" {
		rows, err := models.DB.Query("SELECT codepostal, adresse, date, state, estimatetime, ville, pointrelais, probleme, livre FROM command where idcolis = ?", GetURLID)
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			err := rows.Scan(&packageInfo.CodePostal, &packageInfo.Adresse, &packageInfo.Date, &packageInfo.State, &packageInfo.EstimateTime, &packageInfo.Ville, &packageInfo.PointRelais, &packageInfo.Probleme, &packageInfo.Livre)
			if err != nil {
				panic(err)
			}
		}
		println(packageInfo.Ville)
	}
	tmpl, err := template.ParseFiles("./view/followPackage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, packageInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
