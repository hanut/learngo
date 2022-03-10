package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
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
		data := struct {
			Submissions    url.Values
			RequestDetails struct {
				Method        string
				Path          string
				Header        http.Header
				Host          string
				ContentLength int64
			}
		}{
			Submissions: req.Form,
		}
		data.RequestDetails.Method = req.Method
		data.RequestDetails.Path = req.URL.Path
		data.RequestDetails.Header = req.Header
		data.RequestDetails.Host = req.Host
		data.RequestDetails.ContentLength = req.ContentLength

		tpl.ExecuteTemplate(res, "index.gohtml", data)
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
		AppName: "Forms Microservice 2",
		Env:     "development",
	}

	fmt.Printf("%s (%s) listening at %s\n", app.AppName, app.Env, app.Address())
	http.ListenAndServe(app.Address(), &app)
}
