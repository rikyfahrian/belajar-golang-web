package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type M map[string]interface{}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := M{"nama": "riky"}

		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))

		err := tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := M{"nama": "fahrian"}

		tmpl := template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))

		err := tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server on localhost:8080")
	errr := http.ListenAndServe(":8080", mux)
	log.Fatal(errr)

}
