package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/contact" {
		models.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("./view/contact.html")
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
