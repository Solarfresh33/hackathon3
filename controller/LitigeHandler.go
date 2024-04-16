package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)

func LitigeHandler(w http.ResponseWriter, r *http.Request) {
	var Connected PackageInfo
	cookie, _ := r.Cookie("User")
	if cookie == nil {
		Connected.Connected = false
		println("Connected = false")
	} else {
		Connected.Connected = true
		println("Connected = true")
	
	}

	if r.URL.Path != "/litige" {
		models.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("./view/litige.html")
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
