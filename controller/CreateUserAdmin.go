package controllers

import (
	"encoding/base64"
	"encoding/json"
	models "hackaton/model"
	"html/template"
	"net/http"
)

// boolean admin exist

var adminCreate bool

func CreateUserAdminHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/CreateUser" {
		models.NotFound(w, r)
		return
	}

	session, _ := r.Cookie("User")
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var UserAdmin models.Users
	var CreateAdmin models.Users
	if r.Method == "POST" {
		UserAdmin.Adminusr = r.FormValue("adminusr")
		UserAdmin.Adminpswd = r.PostFormValue("adminpswd")
		if UserAdmin.Adminusr != "" && UserAdmin.Adminpswd != "" {
			CreateAdmin, _ = models.CreateUserAdmin(UserAdmin.Adminusr, UserAdmin.Adminpswd, UserAdmin.Uid, UserAdmin.Admin)
			println("user crée")
			adminCreate = true
			userData := &UserCookie{
				Id:     UserAdmin.Id,
				Pseudo: UserAdmin.Adminusr,
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
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
	if adminCreate {
		println(UserAdmin.Adminusr)
		_, err := models.DB.Exec("DELETE FROM Users WHERE adminusr = ?", "admin")
		if err != nil {
			panic(err)
		}
		adminCreate = false
		println("user supprimé", adminCreate)
	}

	tmpl, err := template.ParseFiles("./view/CreateUser.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, CreateAdmin)
	if err != nil {
		panic(err)
	}
}
