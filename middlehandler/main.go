package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	mux := http.DefaultServeMux

	mux.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Query().Get("id")
		if id != "" {
			beJSON(w, SelectStudent(id))
			return
		}

		beJSON(w, GetStudent())

	})

	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)
	handler = MiddlewareAllowOnlyGet(handler)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = handler

	log.Println("server start on localhost:9000")

	server.ListenAndServe()

}

func beJSON(w http.ResponseWriter, o any) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte("something went wrong"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}
