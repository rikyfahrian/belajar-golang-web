package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case "POST":
	// 		w.Write([]byte("kamu nenye"))
	// 	case "GET":
	// 		w.Write([]byte("kamu bertenye tenye"))
	// 	default:
	// 		http.Error(w, "", http.StatusInternalServerError)
	// 	}
	// })

	//menangkap isi form
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			tmpl := template.Must(template.New("form").ParseFiles("views/form.html"))
			err := tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return

		}

		http.Error(w, "", http.StatusBadRequest)

	})

	mux.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl := template.Must(template.New("result").ParseFiles("views/form.html"))

			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			name := r.FormValue("nama")
			message := r.FormValue("message")

			data := map[string]string{"name": name, "message": message}

			errr := tmpl.Execute(w, data)
			if errr != nil {
				http.Error(w, errr.Error(), http.StatusInternalServerError)

			}
			return
		}

		http.Error(w, "", http.StatusBadRequest)

	})

	log.Println("server on localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err.Error())
	}

}
