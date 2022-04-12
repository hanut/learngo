package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func requestLogger(callback func(w http.ResponseWriter, r *http.Request, p httprouter.Params)) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		st := time.Now()
		callback(w, r, p)
		fmt.Printf("%s -> %s %s [%v]\n", r.RemoteAddr, r.Method, r.URL, time.Since(st))
	}
}

func main() {
	router := httprouter.New()

	router.GET("/", requestLogger(index))
	router.GET("/users", requestLogger(GetUsers))
	router.POST("/users", requestLogger(CreateUsers))

	fmt.Println("HttpRouter based Microservice listening on port 3001...")

	http.ListenAndServe("localhost:3001", router)
}

func index(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.Write([]byte("Server up since " + time.Since(startTime).String()))
}
