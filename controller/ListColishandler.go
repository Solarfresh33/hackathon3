package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)

type PackageList struct {
	ID      int
	Idcolis string
}

func ListColisHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := r.Cookie("User")
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var packageList []PackageList
	rows, err := models.DB.Query("SELECT Id, IdColis FROM Command")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var colis PackageList
		err := rows.Scan(&colis.ID, &colis.Idcolis)
		if err != nil {
			panic(err)
		}
		packageList = append(packageList, colis)
	}
	tmpl, err := template.ParseFiles("./view/listColis.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl.Execute(w, packageList)
}
