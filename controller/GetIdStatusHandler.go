package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)

func GetIDStatusHandler(w http.ResponseWriter, r *http.Request) {
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

	if r.URL.Path != "/up" {
		models.NotFound(w, r)
		return
	}

	getid := r.FormValue("suivie")

	if getid != "" {
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
			if IdColis == getid {
				exist = true
			}
		}
		defer rows.Close()
		if exist {
			http.Redirect(w, r, "/update/"+getid, http.StatusSeeOther)
			return
		}
	}
	tmpl, err := template.ParseFiles("./view/idupdate.html")
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
