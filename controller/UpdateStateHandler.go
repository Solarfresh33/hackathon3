package controllers

import (
	"crypto/md5"
	"encoding/hex"
	models "hackaton/model"
	"html/template"
	"net/http"
	"strings"
)

func UpdateStateHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := r.Cookie("User")
	if session == nil {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}
	var packageInfo PackageInfo
	GetURLID := strings.Split(r.URL.Path, "/")[2]
	packageInfo.Idcolis = GetURLID
	println(packageInfo.Idcolis)
	h := md5.New()
	idStr := GetURLID
	h.Write([]byte(idStr))
	idStr = hex.EncodeToString(h.Sum(nil))
	println("l'url est : ", GetURLID)

	if GetURLID != "" {
		if r.Method == "POST" {
			newState := r.FormValue("etat")
			newpointRelais := r.FormValue("pointRelais")
			livre := r.FormValue("livre")
			probleme := r.FormValue("probleme")
			println("le nouvel etat est : ", newState)
			println("le nouveau point relais est : ", newpointRelais)
			println("le livre est : ", livre)
			println("le probleme est : ", probleme)

			_ = models.UpdatePackageState(idStr, newState, newpointRelais, livre, probleme)
			http.Redirect(w, r, "/id/"+GetURLID, http.StatusSeeOther)
			return
		}
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
