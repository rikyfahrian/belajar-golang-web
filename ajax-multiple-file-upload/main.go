package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("view.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, "only accept method post", http.StatusBadRequest)
			return
		}

		basePath, _ := os.Getwd()
		reader, err := r.MultipartReader()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for {

			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}

			fileLocation := filepath.Join(basePath, "files", part.FileName())
			dst, err := os.Create(fileLocation)
			if dst != nil {
				defer dst.Close()
			}

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if _, err := io.Copy(dst, part); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}

		w.Write([]byte("all files uploaded"))

	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	log.Println("server start on localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err.Error())
	}

}
