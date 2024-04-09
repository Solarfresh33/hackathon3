package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {

	models.CreateDate()

	dates, err := models.GetDate("u")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./view/test.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Dates []models.Date
	}{
		Dates: dates,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
