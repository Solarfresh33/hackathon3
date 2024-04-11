package models

import (
	"html/template"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	tmpl, err := template.ParseFiles("./view/404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
