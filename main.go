package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	authmiddleware "github.com/Colris-Inc/auth/auth-middleware"
)

func main() {
	log.Println("starting auth service")

	//include code for start up service

	http.HandleFunc("/v0/registerUser", registerUser)
	http.HandleFunc("/v0/login", login)
	http.HandleFunc("/v0/logout", logout)
	http.HandleFunc("/v0/userInfo", userInfo)
	http.HandleFunc("/v0/refresh", refresh)
	http.HandleFunc("/v0/updatePassword", updatePassword)
	http.HandleFunc("/v0/updateUser", updateUser)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		w.WriteHeader(401)
		fmt.Fprint(w, `{"message": "Auth Token Missing"}`)
		return
	} else {
		authStatus := authmiddleware.IsAuthenticated(authToken)
		if !authStatus {
			w.WriteHeader(401)
			fmt.Fprint(w, `{"message": "Invalid Auth Token"}`)
			return
		}
	}
	JSONVal, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		fmt.Fprint(w, `{response: `+err.Error()+`}`)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(403)
		fmt.Fprint(w, `{response: POST method expected}`)
		return
	} else {
		response := registerUserHandler(string(JSONVal))
		fmt.Fprint(w, response)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
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
		response := loginHandler(string(JSONVal))
		fmt.Fprint(w, response)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
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
		response := logoutHandler(string(JSONVal))
		fmt.Fprint(w, response)
	}
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	JSONVal, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		fmt.Fprint(w, `{response: `+err.Error()+`}`)
	}
	if r.Method != "GET" {
		w.WriteHeader(403)
		fmt.Fprint(w, `{response: GET method expected}`)
	} else {
		response := userInfoHandler(string(JSONVal))
		fmt.Fprint(w, response)
	}
}

func refresh(w http.ResponseWriter, r *http.Request) {
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
		response := refreshHandler(string(JSONVal))
		fmt.Fprint(w, response)
	}
}

func updatePassword(w http.ResponseWriter, r *http.Request) {
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
		response := updatePasswordHandler(string(JSONVal))
		fmt.Fprint(w, response)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
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
		response := updateUserHandler(string(JSONVal))
		fmt.Fprint(w, response)
	}
}
