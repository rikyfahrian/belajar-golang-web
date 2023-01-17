package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}
type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func (t Info) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said : %s", from, message)
}

var funcMap = template.FuncMap{

	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
	"avg": func(n ...int) int {
		var total = 0
		for _, each := range n {
			total += each
		}
		return total / len(n)
	},
}

const coba = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <p> fakkkkkkkkkkk banh </p>
</body>
</html>`

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name:    "riky",
			Gender:  "Laki laki",
			Hobbies: []string{"skateboarding", "coding"},
			Info:    Info{"bekasi", "markan"},
		}

		tmpl := template.Must(template.ParseFiles("views/view.html"))
		err := tmpl.Execute(w, person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/custom", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("views/custom.html").Funcs(funcMap).ParseFiles("views/custom.html"))

		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	mux.HandleFunc("/function", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name:    "shorthair",
			Gender:  "women",
			Hobbies: []string{"read", "watch anime"},
			Info:    Info{"Indonesia", "Jakarta"},
		}

		tmpl := template.Must(template.ParseFiles("views/pangtien.html"))
		err := tmpl.Execute(w, person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	mux.HandleFunc("/parse", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("nenye").ParseFiles("views/nenye.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	mux.HandleFunc("/nenye", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("kamu").ParseFiles("views/nenye.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	mux.HandleFunc("/coba", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("main-template").Parse(coba))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	mux.HandleFunc("/salah", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/coba", http.StatusTemporaryRedirect)
	})

	fmt.Println("server on running in loclahost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err.Error())
	}

}
