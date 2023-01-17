package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		if !Auth(w, r) {
			return
		}
		if !AllowOnlyGet(w, r) {
			return
		}

		id := r.URL.Query().Get("id")
		if id != "" {
			OutputJSON(w, SelectStudent(id))
			return
		}

		OutputJSON(w, GetStudents())

	})

	log.Println("server start on localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err.Error())
	}

}

func OutputJSON(w http.ResponseWriter, o any) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}
