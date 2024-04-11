package controllers

import (
	
	models "hackaton/model"
	"html/template"
	"net/http"
	"strconv"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var PackageCreate models.UsersPackage
	var CreateColis models.UsersPackage
	if r.Method == "POST" {
		PackageCreate.Adresse = r.FormValue("Adresse")
		PackageCreate.CodePostal, _ = strconv.Atoi(r.FormValue("CodePostal"))
		PackageCreate.Ville = r.FormValue("Ville")
		PackageCreate.PointRelais = r.FormValue("pointRelais")
		if PackageCreate.Adresse != "" && PackageCreate.CodePostal != 0 {
			println("1", PackageCreate.CodePostal, PackageCreate.Adresse)
			CreateColis, _ = models.CreatePackage(PackageCreate.Adresse, PackageCreate.IdColis, PackageCreate.CodePostal, PackageCreate.Date, PackageCreate.State, PackageCreate.EstimateTime, PackageCreate.Ville, PackageCreate.PointRelais)
		}

	}
    println(PackageCreate.Date)
	session, _ := r.Cookie("User")
	if session == nil {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
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
