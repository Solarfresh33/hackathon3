package controllers

import (
	"crypto/md5"
	"encoding/hex"
	models "hackaton/model"
	"html/template"
	"net/http"
)

func GetIDHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
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
		for rows.Next() {
			var IdColis string
			err := rows.Scan(&IdColis)
			if err != nil {
				panic(err)
			}
			if IdColis == idStr {
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

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
