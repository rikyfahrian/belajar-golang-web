package main

import "net/http"

const USERNAME = "kamu"
const PASSWORD = "nanya"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte("wrong username/password"))
		return false
	}

	return true
}

func OnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("something went wrong"))
		return false
	}

	return true

}
