package models

import (
	"net/smtp"
	"strings"
)

func sendMail(to []string, subject, body string) error {
	// Configuration du serveur SMTP
	from := "follow.tracky@gmail.com"
	pass := "jxzi test mgfx kpcs"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587" // Port pour SMTP avec TLS/STARTTLS

	// Message
	message := []byte("To:" + strings.Join(to, ",") + "\r\n" + "Subject: " + subject + "\r\n" + "\r\n" + body + "\r\n")

	// Authentification
	auth := smtp.PlainAuth("", from, pass, smtpHost)

	// Envoi du mail
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil
}

func SendMailReel() {
	rows, err := DB.Query("SELECT idcolis FROM command")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var idcolis string
	for rows.Next() {
		err := rows.Scan(&idcolis)
		if err != nil {
			panic(err)
		}
	}
	userPackage := GetEmailandNAme(idcolis)

	to := []string{userPackage.Email}
	subject := "Votre colis Tracky !"
	body := "Votre colis a été mis à jour.\n Veuillez le consultez avec le lien ci-contre  : https://groupe5.etudiants.ynov-bordeaux.com/id/" + idcolis  
	err = sendMail(to, subject, body)
	if err != nil {
		panic(err)
	}
}
