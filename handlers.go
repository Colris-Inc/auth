package main

import (
	"encoding/json"
	"log"

	authmiddleware "github.com/Colris-Inc/auth/auth-middleware"
	"github.com/Colris-Inc/auth/db"
	"github.com/Colris-Inc/auth/encryption"
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

type updatePasswordData struct {
	User            string `json:"user"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

func registerUserHandler(jsonString string) string {
	var response string
	log.Println("registering user")
	registerData := &userRegistrationData{}
	json.Unmarshal([]byte(jsonString), &registerData)

	encryptedPassword := encryption.EncryptString(registerData.InitialPassword)

	err := db.RegisterUser(registerData.UserName, registerData.FirstName, registerData.LastName, registerData.MobileNumber,
		registerData.Email, encryptedPassword, registerData.Roles)

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
	inputStruct := &updatePasswordData{}
	json.Unmarshal([]byte(jsonString), &inputStruct)

	userStatus := authmiddleware.ValidatePassword(inputStruct.User, inputStruct.CurrentPassword)

	if !userStatus {
		response = `{"message": "username or password incorrect"}`
		return response
	}

	err := db.UpdatePassword(inputStruct.User, inputStruct.CurrentPassword, inputStruct.NewPassword)
	if err != nil {
		log.Println("[Error Occured]: " + err.Error())
		response = `{"message": "` + err.Error() + `"}`
		return response
	}

	response = `{"message": "password updated successfully"}`

	return response
}

func updateUserHandler(jsonString string) string {
	var response string

	return response
}
