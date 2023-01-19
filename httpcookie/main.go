package main

import (
	"log"
	"net/http"
	"time"

	gubrak "github.com/novalagung/gubrak/v2"
)

type M map[string]any

var cookieName = "CookieData"

func main() {

	mux := http.DefaultServeMux

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookieName := "CookieData"

		c := &http.Cookie{}

		if isit, _ := r.Cookie(cookieName); isit != nil {
			c = isit
		}

		if c.Value == "" {
			c := &http.Cookie{}
			c.Name = cookieName
			c.Value = gubrak.RandomString(32)
			c.Expires = time.Now().Add(5 * time.Minute)

			http.SetCookie(w, c)
			w.Write([]byte(c.Value))
		}

		w.Write([]byte(c.Value))

	})

	mux.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {

		c := &http.Cookie{}
		c.Name = cookieName
		c.Expires = time.Unix(0, 0)
		c.MaxAge = -1
		http.SetCookie(w, c)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	log.Println("server on localhost:8080")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Println(err.Error())
	}

}
