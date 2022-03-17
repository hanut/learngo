package main

import (
	"fmt"
	"net/http"
	"time"
)

type DaytimeServer bool

func (ds DaytimeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Print some the request details
	fmt.Println("Headers", r.Header)
	fmt.Println("Request Path", r.URL.Path)
	fmt.Println("Query", r.URL.RawQuery)

	// Send back the date and time
	w.Write([]byte(time.Now().String()))
}

func main() {
	fmt.Println("Http Server in Golang")

	// Create a new Daytime Server
	var ds DaytimeServer

	fmt.Println("Server is starting at localhost:8080")

	// To create an http server we use the net/http package
	http.ListenAndServe("localhost:8080", ds)

}
