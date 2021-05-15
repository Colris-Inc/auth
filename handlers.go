package main

import (
	"encoding/json"
	"log"

	"github.com/Colris-Inc/auth/db"
)

type userRegistrationData struct {
	UserName        string   `json:"userName"`
	FirstName       string   `json:"firstName"`
	LastName        string   `json:"lastName"`
	MobileNumber    int      `json:"mobileNumber"`
	Email           string   `json:"email"`
	InitialPassword string   `json:"initialPassword"`
	Roles           []string `json:"roles"`
}

func registerUserHandler(jsonString string) string {
	var response string
	log.Println("registering user")
	registerData := &userRegistrationData{}
	json.Unmarshal([]byte(jsonString), &registerData)

	err := db.RegisterUser(registerData.UserName, registerData.FirstName, registerData.LastName, registerData.MobileNumber,
		registerData.Email, registerData.InitialPassword, registerData.Roles)

	if err != nil {
		log.Println("[Error Occured]: " + err.Error())
		response = `{"message": "` + err.Error() + `"}`
		return response
	}
	response = `{"message": "successfully registered user"}`
	log.Println("user registration completed")
	return response
}

func loginHandler(jsonString string) string {
	var response string

	return response
}

func logoutHandler(jsonString string) string {
	var response string

	return response
}

func userInfoHandler(jsonString string) string {
	var response string

	return response
}

func refreshHandler(jsonString string) string {
	var response string

	return response
}

func updatePasswordHandler(jsonString string) string {
	var response string

	return response
}

func updateUserHandler(jsonString string) string {
	var response string

	return response
}
