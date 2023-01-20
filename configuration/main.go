package main

import (
	"configg/conf"
	"fmt"
	"log"
	"net/http"
)

type CustomMux struct {
	http.ServeMux
}

func (c CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if conf.Configuration().Log.Verbose {
		log.Println("Incoming request form", r.Host, "accesing", r.URL.String())

	}

	c.ServeMux.ServeHTTP(w, r)
}

func main() {

	router := new(CustomMux)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("hello ini halaman home"))

	})

	router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("ini halaman about"))

	})

	server := new(http.Server)
	server.Handler = router
	server.ReadTimeout = conf.Configuration().Server.ReadTimeOut
	server.WriteTimeout = conf.Configuration().Server.WriteTimeOut
	server.Addr = fmt.Sprintf(":%d", conf.Configuration().Server.Port)

	if conf.Configuration().Log.Verbose {
		log.Println("server start on localhost:9000")
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}
