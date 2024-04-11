package controllers

import (
	"crypto/md5"
	"encoding/hex"
	models "hackaton/model"
	"html/template"
	"net/http"
)

func GetIDStatusHandler(w http.ResponseWriter, r *http.Request) {
	
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
		h := md5.New()
		idStr := getid
		h.Write([]byte(idStr))
		idStr = hex.EncodeToString(h.Sum(nil))
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
			if IdColis == idStr {
				exist = true
			}
		}
		defer rows.Close()
		if exist {
			http.Redirect(w, r, "/update/"+getid, http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl, err := template.ParseFiles("./view/idupdate.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
