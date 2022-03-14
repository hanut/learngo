package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

	func(res http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("received form")
		for k, v := range req.Form {
			fmt.Println(k, v)
		}
		tpl.ExecuteTemplate(res, "index.gohtml", req.Form)
	}(res, req)

	fmt.Printf("%s -> %s %s [%v]\n", req.RemoteAddr, req.Method, req.URL, time.Since(st))

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	app := myHttpServer{
		Host:    "localhost",
		Port:    3001,
		AppName: "Forms Microservice 1",
		Env:     "development",
	}

	fmt.Printf("%s (%s) listening at %s\n", app.AppName, app.Env, app.Address())
	http.ListenAndServe(app.Address(), &app)
}
