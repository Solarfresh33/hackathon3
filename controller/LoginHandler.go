package controllers

import (
	"encoding/base64"
	"encoding/json"
	"sync"

	// "fmt"
	models "hackaton/model"
	"html/template"

	// "log"
	"net/http"
)

var adminCreated bool
var adminMutex sync.Mutex

type UserCookie struct {
	Id     int    `json:"id"`
	Pseudo string `json:"pseudo"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var Connected PackageInfo
	cookie, _ := r.Cookie("User")
	if cookie == nil {
		Connected.Connected = false
		println("Connected = false")
	} else {
		Connected.Connected = true
		println("Connected = true")

	}

	if r.URL.Path != "/login" {
		models.NotFound(w, r)
		return
	}

	var id int
	// var uid models.UsersPackage
	adminMutex.Lock()
	defer adminMutex.Unlock()

	if !adminCreated {
		models.FirstUserAdmin()
		adminCreated = true
	}
	// rows, err := models.DB.Query("SELECT adminusr, adminpswd FROM Users WHERE uid=?", uid)
	// if err != nil {
	// 	panic(err)
	// }
	// var adminuser string
	// var adminpassword string
	// for rows.Next() {

	// 	err := rows.Scan(&adminuser, &adminpassword)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	if r.Method == "POST" {
		pseudo := r.FormValue("pseudo")
		password := r.FormValue("password")

		if pseudo != "" && password != "" {
			existaccount, psswd, _ := models.ExistAccount(pseudo)
			if existaccount && psswd == password {
				userData := &UserCookie{
					Id:     id,
					Pseudo: pseudo,
				}

				userBytes, err := json.Marshal(userData)
				if err != nil {
					panic(err)
				}

				http.SetCookie(w, &http.Cookie{
					Name:     "User",
					Value:    base64.URLEncoding.EncodeToString(userBytes),
					HttpOnly: true,
					MaxAge:   604800,
				})

				http.Redirect(w, r, "/admin", http.StatusSeeOther)
				return
			}

		}
	}

	tmpl, err := template.ParseFiles("./view/login.html")
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
