package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
	"strings"
)

func UpdateStateHandler(w http.ResponseWriter, r *http.Request) {
	var packageInfo PackageInfo

	cookie, _ := r.Cookie("User")
	if cookie == nil {
		packageInfo.Connected = false
		println("Connected = false")
	} else {
		packageInfo.Connected = true
		println("Connected = true")
	
	}
	session, _ := r.Cookie("User")
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	GetURLID := strings.Split(r.URL.Path, "/")[2]
	packageInfo.Idcolis = GetURLID
	println(packageInfo.Idcolis)

	if r.URL.Path != "/update/"+GetURLID {
		models.NotFound(w, r)
		return
	}

	if GetURLID != "" {
		rows, err := models.DB.Query("SELECT idcolis FROM command")
		if err != nil {
			panic(err)
		}
		var exist bool
		for rows.Next() {
			var IdColis string
			err := rows.Scan(&IdColis)
			if err != nil {
				panic(err)
			}
			if IdColis == GetURLID {
				exist = true
			}
		}
		defer rows.Close()
		if !exist {
			http.Redirect(w, r, "/scan", http.StatusSeeOther)
			return
		}
	}
	if r.Method == "POST" {
		newState := r.FormValue("etat")
		newpointRelais := r.FormValue("pointRelais")
		livre := r.FormValue("livre")
		probleme := r.FormValue("probleme")
		_ = models.UpdatePackageState(GetURLID, newState, newpointRelais, livre, probleme)
		models.SendMailReel()
		http.Redirect(w, r, "/id/"+GetURLID, http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./view/state.html")
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
