package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
	"strconv"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var CreateColis models.UsersPackage

	cookie, _ := r.Cookie("User")
	if cookie == nil {
		CreateColis.Connected = false
		println("Connected = false")
	} else {
		CreateColis.Connected = true
		println("Connected = true")
	
	}
	if r.URL.Path != "/create" {
		models.NotFound(w, r)
		return
	}

	var PackageCreate models.UsersPackage
	if r.Method == "POST" {
		PackageCreate.Adresse = r.FormValue("Adresse")
		PackageCreate.CodePostal, _ = strconv.Atoi(r.FormValue("CodePostal"))
		PackageCreate.Ville = r.FormValue("Ville")
		PackageCreate.PointRelais = r.FormValue("pointRelais")
		PackageCreate.Email = r.FormValue("Email")
		PackageCreate.Name = r.FormValue("Name")
		if PackageCreate.Adresse != "" && PackageCreate.CodePostal != 0 {
			CreateColis, _ = models.CreatePackage(PackageCreate.Adresse, PackageCreate.IdColis, PackageCreate.Email, PackageCreate.Name, PackageCreate.CodePostal, PackageCreate.Date, PackageCreate.State, PackageCreate.EstimateTime, PackageCreate.Ville, PackageCreate.PointRelais)
			models.SendMailReel()
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
	

		}

	}
	session, _ := r.Cookie("User")
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./view/index.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, CreateColis)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
