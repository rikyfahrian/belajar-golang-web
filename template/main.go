package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandle)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err.Error())
	}

}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("views", "index.html")
	parse, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"title": "home",
		"name":  "riky",
	}

	err = parse.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
