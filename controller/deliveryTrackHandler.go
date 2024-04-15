package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)

type PackageList struct {
	ID           int
	Idcolis      string
	CodePostal   string
	Adresse      string
	Name         string
	State        string
	Date         string
	EstimateTime string
	Ville        string
	PointRelais  string
	Probleme     string
	Livre        string
}

func DeliveryTrackHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := r.Cookie("User")
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var packageList []PackageList
	rows, err := models.DB.Query("SELECT Id, IdColis,CodePostal,Adresse,Name,State,Date,EstimateTime,ville,Livre,Probleme,PointRelais FROM Command")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var colis PackageList
		err := rows.Scan(&colis.ID, &colis.Idcolis, &colis.CodePostal, &colis.Adresse, &colis.Name, &colis.State, &colis.Date, &colis.EstimateTime, &colis.Ville, &colis.Livre, &colis.Probleme, &colis.PointRelais)
		if err != nil {
			panic(err)
		}
		packageList = append(packageList, colis)
	}
	tmpl, err := template.ParseFiles("./view/deliveryTrack.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, packageList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
