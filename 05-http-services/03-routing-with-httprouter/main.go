package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Routing using HttpRouter")

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("Sending back datetime")
		fmt.Println("Http Method Used:", r.Method)
		w.Write([]byte(time.Now().String()))
	})

	router.GET("/users", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("Sending back users")
		fmt.Println("Http Method Used:", r.Method)
		w.Write([]byte("Alice, Bob, Charlie,David,Englebert,Henry"))
	})

	router.GET("/pets", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("Sending back users")
		fmt.Println("Http Method Used:", r.Method)
		w.Write([]byte("Tommy, Moti, Sheena, Billu, Bhura, Kalu"))
	})

	// Accessing and using the query parameters from a request
	router.GET("/search/users", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("Search Query:", r.URL.Query().Get("name"))

		msg := "Searched for user:\t" + r.URL.Query().Get("name")

		w.Write([]byte(msg))
	})

	// Parametrized Route which takes a user name as a path parameter
	router.GET("/users/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("Fetching user by name", p.ByName("name"))

		msg := "Requested User:\t" + p.ByName("name")
		w.Write([]byte(msg))
	})

	fmt.Println("HttpRouter based service listening on localhost:8082 ...")
	http.ListenAndServe("localhost:8082", router)
}
