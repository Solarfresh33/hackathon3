package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var PackageCreate  models.UsersPackage 
	var CreateColis models.UsersPackage
	if r.Method == "POST" {
		PackageCreate.Adresse = r.FormValue("Adresse")
		PackageCreate.CodePostal ,_= strconv.Atoi(r.FormValue("CodePostal"))
		println("1",PackageCreate.CodePostal,PackageCreate.Adresse)
		CreateColis, _ = models.CreatePackage(PackageCreate.Adresse, PackageCreate.IdUser, PackageCreate.CodePostal)
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
