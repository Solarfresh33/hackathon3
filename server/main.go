package main

import (
	"fmt"
	controllers "hackaton/controller"
	models "hackaton/model"
	"net/http"
)

func main() {

	models.OpenDateBase()
	models.CreateDBUsers()
	models.CreateDBCommand()
	defer models.DB.Close()

	cssHandler := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", cssHandler))
	http.HandleFunc("/create", controllers.CreateHandler)
	http.HandleFunc("/", controllers.GetIDHandler)
	http.HandleFunc("/id/", controllers.InfoPackageHandler)
	http.HandleFunc("/contact", controllers.ContactHandler)
	http.HandleFunc("/FAQ", controllers.FAQHandler)
	http.HandleFunc("/update/", controllers.UpdateStateHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/admin", controllers.MenuboardAdminHandler)
	http.HandleFunc("/CreateUser", controllers.CreateUserAdminHandler)
	http.HandleFunc("/deliveryTrack", controllers.DeliveryTrackHandler)
	http.HandleFunc("/scan", controllers.QRCodeHandler)
	http.HandleFunc("/up", controllers.GetIDStatusHandler)
	http.HandleFunc("/deco", controllers.DecoHandler)
	http.HandleFunc("/litige", controllers.LitigeHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
