package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("view.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	mux.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			payload := struct {
				Name string `json:"name"`

				Age    int    `json:"age"`
				Gender string `json:"gender"`
			}{}

			err := decoder.Decode(&payload)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			message := fmt.Sprintf("hello my name is %s, i'm %d years old %s", payload.Name, payload.Age, payload.Gender)

			w.Write([]byte(message))
			return

		}

		http.Error(w, "", http.StatusBadRequest)
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	log.Println("server start on localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err.Error())
	}

}
