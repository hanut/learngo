package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", Index)
	r.Get("/users", GetUsers)
	r.Post("/users", CreateUsers)

	fmt.Println("Chi Router based Microservice listening on port 3001...")

	http.ListenAndServe("localhost:3001", r)

}

func Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Server up since " + time.Since(startTime).String()))
}
