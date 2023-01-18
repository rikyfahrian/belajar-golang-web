package main

import "net/http"

const USERNAME = "kamu"
const PASSWORD = "nanya"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("something went wrong"))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte("wrong username/password"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			w.Write([]byte("something went wrong"))
			return
		}

		next.ServeHTTP(w, r)

	})
}
