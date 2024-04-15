package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)

func MenuboardAdminHandler(w http.ResponseWriter, r *http.Request) {
	var Connected PackageInfo
	cookie, _ := r.Cookie("User")
	if cookie == nil {
		Connected.Connected = false
		println("Connected = false")
	} else {
		Connected.Connected = true
		println("Connected = true")
	
	}
	session, _ := r.Cookie("User")
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.URL.Path != "/admin" {
		models.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("./view/menuboardAdmin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, Connected)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
