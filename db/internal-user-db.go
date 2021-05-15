package db

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type claimsJSON struct {
	Claims []string `json:"claims"`
}

//createPasswordResetDate returns a date with 3 months added to current date for password reset
func createPasswordResetDate() string {
	var response string
	var month int

	currentTime := time.Now().UTC()
	resetTime := currentTime.AddDate(0, 3, 0)
	switch resetTime.Month().String() {
	case "January":
		month = 1
	case "February":
		month = 2
	case "March":
		month = 3
	case "April":
		month = 4
	case "May":
		month = 5
	case "June":
		month = 6
	case "July":
		month = 7
	case "August":
		month = 8
	case "September":
		month = 9
	case "October":
		month = 10
	case "November":
		month = 11
	case "December":
		month = 12
	}
	response = strconv.Itoa(resetTime.Day()) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(resetTime.Year())

	return response
}

//RegisterUser registers a new user into the application
func RegisterUser(userName string, firstName string, lastName string, mobileNumber int,
	email string, password string, roles []string) error {
	db, err := connectDB()
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return err
	}
	defer db.Close()
	claimsStruct := &claimsJSON{}
	claimsStruct.Claims = roles
	encodedClaims, _ := json.Marshal(claimsStruct)
	claims := string(encodedClaims)
	query := "insert into internal_users (user_name, first_name, last_name, email, claims," +
		"password, password_reset, mobile_number, reset_on_login) vales ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	passwordResetDate := createPasswordResetDate()
	_, err = db.Exec(query, userName, firstName, lastName, email, claims, password, passwordResetDate,
		mobileNumber, true)
	return err
}
