package controllers

import (
	// "crypto/md5"
	// "encoding/hex"
	models "hackaton/model"
	"html/template"
	"net/http"
	// "strings"
)

func QRCodeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/scan" {
		models.NotFound(w, r)
		return
	}

	session, _ := r.Cookie("User")
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./view/scan.html")
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
