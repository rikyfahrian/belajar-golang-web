package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		tmpl := template.Must(template.ParseFiles("view.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

	})

	mux.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			nama := r.FormValue("nama")
			umur := r.FormValue("umur")

			data := []struct {
				Nama string
				Age  string
			}{
				{nama, umur},
			}

			hasil, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(hasil)

		}

		http.Error(w, "", http.StatusBadRequest)
	})

	log.Println("server start on localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err.Error())
	}
}
