package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type User struct {
	Id   string
	Name string
}

var tpls *template.Template
var users []User

func init() {
	// Setup users
	users = make([]User, 3, 3)
	users[0] = User{Id: "10001", Name: "Alice"}
	users[1] = User{Id: "10002", Name: "Bob"}
	users[2] = User{Id: "10003", Name: "Charlie"}

	// Setup template
	t, err := template.ParseFiles("index.gohtml", "results.gohtml")
	if err != nil {
		panic(err)
	}
	tpls = t
}

func main() {
	http.DefaultServeMux.HandleFunc("/set-cookie", WithLogger(handleSetCookie))
	http.DefaultServeMux.HandleFunc("/clear-cookie", WithLogger(handleRemoveCookie))
	http.DefaultServeMux.HandleFunc("/results", WithLogger(serveResults))
	http.DefaultServeMux.HandleFunc("/specific", WithLogger(serveSpecificPath))
	http.DefaultServeMux.HandleFunc("/", WithLogger(serveIndex))

	fmt.Println("App listening on localhost:3000...")
	http.ListenAndServe("localhost:3000", http.DefaultServeMux)
}

func WithLogger(rh func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// store the start time
		ts := time.Now()
		// Call the route handler callback
		rh(w, r)
		fmt.Printf("[%s] %s (%s)\n", r.Method, r.URL, time.Since(ts).String())
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	var qe string
	if r.URL.Query().Has("error") {
		qe = r.URL.Query().Get("error")
	}
	var cs bool
	_, err := r.Cookie("myCookie")
	if err == nil {
		cs = true
	}
	data := struct {
		Users  []User
		Error  string
		Cookie bool
	}{Users: users, Error: qe, Cookie: cs}
	tpls.ExecuteTemplate(w, "index.gohtml", data)
}

func serveResults(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("myCookie")
	if err != nil {
		http.Redirect(w, r, "/?error=nocookie", 303)
		return
	}
	data := struct {
		CookieCount int
		Cookie      string
	}{
		CookieCount: len(r.Cookies()),
		Cookie:      c.String(),
	}
	tpls.ExecuteTemplate(w, "results.gohtml", data)
}

func serveSpecificPath(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("path-cookie")
	if err != nil {
		http.Redirect(w, r, "/?error=nopathcookie", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println("Path Cookie", c)
	data := struct {
		CookieCount int
		Cookie      string
	}{
		CookieCount: len(r.Cookies()),
		Cookie:      c.String(),
	}
	tpls.ExecuteTemplate(w, "results.gohtml", data)
}

func handleSetCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("myCookie")
	if err == nil {
		http.Redirect(w, r, "/?error=hascookie", http.StatusTemporaryRedirect)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Println("Error:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	userId := r.PostFormValue("userId")
	if userId == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fi := -1
	for i := range users {
		if users[i].Id == userId {
			fi = i
			break
		}
	}
	if fi == -1 {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	cv := base64.StdEncoding.Strict().EncodeToString([]byte(users[fi].Id + "|" + users[fi].Name))
	c = &http.Cookie{
		Name:  "myCookie",
		Value: cv,
	}
	cspecific := &http.Cookie{
		Name:  "path-cookie",
		Value: "A valuable secret for the path",
		Path:  "/specific",
	}
	http.SetCookie(w, c)
	http.SetCookie(w, cspecific)
	http.Redirect(w, r, "/", 307)
}

func handleRemoveCookie(w http.ResponseWriter, r *http.Request) {
	c, ce := r.Cookie("myCookie")
	var loc string = "/"

	if ce != nil {
		loc += "?error=nocookie"
	}

	c = &http.Cookie{
		Name:   "myCookie",
		MaxAge: -1,
	}
	csp := &http.Cookie{
		Name:   "path-cookie",
		Path:   "/specific",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.SetCookie(w, csp)
	http.Redirect(w, r, loc, 303)
}
