package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("starting auth service")

	//include code for start up service

	http.HandleFunc("/v0/registerUser", registerUser)
	//using register user as dummy for now
	http.HandleFunc("/v0/login", registerUser)
	http.HandleFunc("/v0/loout", registerUser)
	http.HandleFunc("/v0/userInfo", registerUser)
	http.HandleFunc("/v0/refresh", registerUser)
	http.HandleFunc("/v0/updatePassword", registerUser)
	http.HandleFunc("/v0/updateUser", registerUser)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	JSONVal, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		fmt.Fprint(w, `{response: `+err.Error()+`}`)
	}
	if r.Method != "POST" {
		w.WriteHeader(403)
		fmt.Fprint(w, `{response: POST method expected}`)
	} else {
		response := registerUserHandler(string(JSONVal))
		fmt.Fprint(w, response)
	}
}
