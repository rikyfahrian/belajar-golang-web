package main

import (
	"encoding/json"
	"net/http"
)

// student
type Student struct {
	Id    string
	Name  string
	Grade int32
}

var students = []*Student{}

func init() {
	students = append(students, &Student{"111", "riky", 3})
	students = append(students, &Student{"112", "riky", 1})
	students = append(students, &Student{"113", "riky", 2})

}

func selectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil

}

func getStudent() []*Student {
	return students
}

// middleware
const user = "kamu"
const pass = "nanya"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("something went wrong"))
		}

		isValid := (username == user) && (password == pass)
		if !isValid {
			w.Write([]byte("wrong username/password"))
			return
		}

		next.ServeHTTP(w, r)

	})
}

func onlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			w.Write([]byte("method must be get"))
			return
		}

		next.ServeHTTP(w, r)

	})
}

type customMux struct {
	http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func (c *customMux) regis(next func(http.Handler) http.Handler) {

	c.middlewares = append(c.middlewares, next)

}

func (c *customMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}

	current.ServeHTTP(w, r)

}

func main() {

	mux := new(customMux)

	mux.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Query().Get("id")
		if id != "" {
			beJSON(w, selectStudent(id))
			return
		}

		beJSON(w, getStudent())

	})

	mux.regis(Auth)
	mux.regis(onlyGet)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = mux

	server.ListenAndServe()

}

func beJSON(w http.ResponseWriter, o any) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}
