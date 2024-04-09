package controllers

import (
	"html/template"
	"net/http"
)

func Home2Handler(w http.ResponseWriter, r *http.Request) {

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
