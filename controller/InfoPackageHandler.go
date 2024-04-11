package controllers

import (
	"crypto/md5"
	"encoding/hex"
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
}

func InfoPackageHandler(w http.ResponseWriter, r *http.Request) {
	var packageInfo PackageInfo
	GetURLID := strings.Split(r.URL.Path, "/")[2]

	packageInfo.Idcolis = GetURLID
	h := md5.New()
	idStr := GetURLID
	h.Write([]byte(idStr))
	idStr = hex.EncodeToString(h.Sum(nil))
	println("l'url est : ", GetURLID)

	if GetURLID != "" {
		rows, err := models.DB.Query("SELECT codepostal, adresse, date, state, estimatetime, ville, pointrelais, probleme, livre FROM command where idcolis = ?", idStr)
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			err := rows.Scan(&packageInfo.CodePostal, &packageInfo.Adresse, &packageInfo.Date, &packageInfo.State, &packageInfo.EstimateTime, &packageInfo.Ville, &packageInfo.PointRelais, &packageInfo.Probleme, &packageInfo.Livre)
			if err != nil {
				panic(err)
			}
		}
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
