package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)

func LitigeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/litige" {
		models.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("./view/litige.html")
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
