package authmiddleware

import (
	"log"

	"github.com/Colris-Inc/auth/db"
	"github.com/Colris-Inc/auth/encryption"
)

//IsAuthenticated validates current auth status by comparing auth token with DB
func IsAuthenticated(token string) bool {
	response := false
	log.Println("checking user auth")

	log.Println("auth check completed")
	return response
}

//ValidatePassword validates the current password against the database
func ValidatePassword(identifier string, password string) bool {
	response := false
	log.Println("checking password")
	encryptedPassword, err := db.FetchPassword(identifier)
	if err != nil {
		log.Println("[Error Occured] : " + err.Error())
		return false
	}
	decryptedPassword := encryption.DecryptString(encryptedPassword)
	if password == decryptedPassword {
		response = true
	} else {
		response = false
	}
	log.Println("password check completed")
	return response
}
