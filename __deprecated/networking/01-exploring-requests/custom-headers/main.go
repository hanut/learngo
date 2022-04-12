package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type myHttpServer struct {
	Host    string
	Port    uint16
	AppName string
	Env     string
}

func (mhs *myHttpServer) Address() string {
	return fmt.Sprintf("%s:%d", mhs.Host, mhs.Port)
}

func (mhs *myHttpServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	st := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)
	go func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Custom-Header-1", "A cool header with secret stuff")
		res.Header().Set("Custom-Header-2", "Another header with app specific stuff")

		if strings.Contains(req.URL.Path, "plaintext") {
			res.Header().Set("Content-Type", "text/plain; charset=utf-8")
		} else {
			res.Header().Set("Content-Type", "text/html; charset=utf-8")
		}
		fmt.Fprintln(res, "<h2>You can send back html from here</h2>")
		wg.Done()
	}(res, req)
	wg.Wait()

	fmt.Printf("%s -> %s %s [%v]\n", req.RemoteAddr, req.Method, req.URL, time.Since(st))
}

func main() {
	app := myHttpServer{
		Host:    "localhost",
		Port:    3001,
		AppName: "Forms Microservice 3",
		Env:     "development",
	}

	fmt.Printf("%s (%s) listening at %s\n", app.AppName, app.Env, app.Address())
	http.ListenAndServe(app.Address(), &app)
}
