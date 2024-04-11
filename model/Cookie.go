package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserCookie struct {
	Id     int    `json:"id"`
	Pseudo string `json:"pseudo"`
}

func GetCookie(r *http.Request) (UserCookie, bool) {
	var userData UserCookie
	cookie, err := r.Cookie("User")
	if err != nil {
		fmt.Println(err)
		return userData, false
	}

	userBytes, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		fmt.Println(err)
		return userData, false
	}

	err = json.Unmarshal(userBytes, &userData)
	if err != nil {
		fmt.Println(err)
		return UserCookie{}, false
	}

	return userData, true
}
