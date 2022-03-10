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

// Method for myHttpServer instances that combines the host and port
// into a uri string and returns it
func (mhs *myHttpServer) Address() string {
	return fmt.Sprintf("%s:%d", mhs.Host, mhs.Port)
}

// myHttpServer instance
var app myHttpServer

// Method for myHttpServer that handles the requests
func (mhs *myHttpServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	st := time.Now()

	switch req.URL.Path {
	case "/":
		tpl.ExecuteTemplate(res, "login.gohtml", nil)
	case "/login":
		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		u := req.FormValue("username")
		ps := req.FormValue("pass")
		if !isAuth(&u, &ps) {
			http.Error(res, "Unauthorized", 401)
			break
		}
		http.Redirect(res, req, "/dashboard?user="+u+"&pass="+ps, 302)
	case "/logout":
		// Simple redirection to home
		http.Redirect(res, req, "/", 301)
	case "/dashboard":
		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
			break
		}
		u := req.FormValue("user")
		ps := req.FormValue("pass")
		err = tpl.ExecuteTemplate(res, "dashboard.gohtml", struct {
			User     string
			Password string
		}{
			User:     u,
			Password: ps,
		})
		if err != nil {
			http.Error(res, err.Error(), 500)
			break
		}
	default:
		err := tpl.ExecuteTemplate(res, "404.gohtml", nil)
		if err != nil {
			http.Error(res, err.Error(), 500)
			break
		}
	}
	fmt.Printf("%s -> %s %s [%v]\n", req.RemoteAddr, req.Method, req.URL, time.Since(st))
}

// Simple hardcoded authentication function that checks
// if the given username and password match the configured
// values
func isAuth(u *string, ps *string) bool {
	if (*u) == "admin" && (*ps) == "test@123" {
		return true
	}
	return false
}

// template instance that holds all the go templates
var tpl *template.Template

// Initialization function that sets up the app server instance
// and loads the required go templates
func init() {
	app = myHttpServer{
		Host:    "localhost",
		Port:    3001,
		AppName: "Routing Microservice 1",
		Env:     "development",
	}
	tpl = template.Must(template.ParseFiles("views/login.gohtml", "views/dashboard.gohtml", "views/404.gohtml"))
}

// Main application function that starts the server
func main() {
	fmt.Printf("%s (%s) listening at %s\n", app.AppName, app.Env, app.Address())
	http.ListenAndServe(app.Address(), &app)
}
