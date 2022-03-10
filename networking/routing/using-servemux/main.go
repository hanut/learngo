package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyApiServer struct {
	AppName string
}

// Method for MyApiServer that handles the requests
func (mhs *MyApiServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	st := time.Now()

	res.Write([]byte(fmt.Sprintf("Response from service %s", mhs.AppName)))

	fmt.Printf("%s -> %s %s [%v]\n", req.RemoteAddr, req.Method, req.URL, time.Since(st))
}

// MyApiServer instance
var apiV1 MyApiServer
var apiV2 MyApiServer

// Initialization function that sets up the app server instance
// and loads the required go templates
func init() {
	apiV1 = MyApiServer{
		AppName: "ServeMux Api v1",
	}
	apiV2 = MyApiServer{
		AppName: "ServeMux Api v2",
	}
}

// Main application function that starts the server
func main() {
	fmt.Println("ServeMux versioned api listening at 3001")
	// Create new mux instance
	mux := http.NewServeMux()

	// Setup handlers for different versions of the api
	mux.Handle("/v1/", &apiV1)
	mux.Handle("/v2/", &apiV2)
	// The next handler doesnt catch anything except the /v3
	// because of a missing trailing slash
	mux.Handle("/v3", &apiV2)

	http.ListenAndServe("localhost:3001", mux)
}
