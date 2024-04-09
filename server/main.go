package main

import (
	"fmt"
	controllers "hackaton/controller"
	models "hackaton/model"
	"net/http"
)

func main() {

	models.OpenDateBase()
	// models.DeleteDB()
	models.CreateDBCommand()
	defer models.DB.Close()

	cssHandler := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", cssHandler))
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/home2", controllers.Home2Handler)
	http.HandleFunc("/test", controllers.Test)
	http.HandleFunc("/contact", controllers.ContactHandler)
	http.HandleFunc("/FAQ", controllers.FAQHandler)
	http.HandleFunc("/404", controllers.NotFoundHandler)


	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
