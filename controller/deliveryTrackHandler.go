package controllers

import (
	"html/template"
	"net/http"
)

func DeliveryTrackHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./view/deliveryTrack.html")
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
