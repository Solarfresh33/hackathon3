package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)


func GetIDHandler(w http.ResponseWriter, r *http.Request) {
	var Connected PackageInfo

	cookie, _ := r.Cookie("User")
	if cookie == nil {
		Connected.Connected = false
		println("Connected = false")
	} else {
		Connected.Connected = true
		println("Connected = true")
	
	}

	if r.URL.Path != "/" {
		models.NotFound(w, r)
		return
	}

	getid := r.FormValue("suivie")

	if getid != "" {
		rows, err := models.DB.Query("SELECT idcolis FROM command")
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var IdColis string
			err := rows.Scan(&IdColis)
			if err != nil {
				panic(err)
			}
			if IdColis == getid {
				http.Redirect(w, r, "/id/"+getid, http.StatusSeeOther)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl, err := template.ParseFiles("./view/index2.html")
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
