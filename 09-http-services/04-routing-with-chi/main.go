package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Routing with Chi")

	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Sending back datetime")
		fmt.Println("Http Method Used:", r.Method)
		w.Write([]byte(time.Now().String()))
	})

	fmt.Println("Chi Router Service listening on localhost:8083")
	http.ListenAndServe("localhost:8083", r)
}
