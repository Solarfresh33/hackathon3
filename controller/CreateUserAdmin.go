package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)

// boolean admin exist

var adminCreate bool

func CreateUserAdminHandler(w http.ResponseWriter, r *http.Request) {
	var CreateAdmin models.Users
	cookie, _ := r.Cookie("User")
	if cookie == nil {
		CreateAdmin.Connected = false
		println("Connected = false")
	} else {
		CreateAdmin.Connected = true
		println("Connected = true")
	
	}
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
	if r.Method == "POST" {
		UserAdmin.Adminusr = r.FormValue("adminusr")
		UserAdmin.Adminpswd = r.PostFormValue("adminpswd")
		if UserAdmin.Adminusr != "" && UserAdmin.Adminpswd != "" {
			CreateAdmin, _ = models.CreateUserAdmin(UserAdmin.Adminusr, UserAdmin.Adminpswd, UserAdmin.Uid, UserAdmin.Admin)
			adminCreate = true
		}
	}
	if adminCreate {
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
