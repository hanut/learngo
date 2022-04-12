package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

var tpls *template.Template

func init() {
	t, err := template.ParseFiles("index.gohtml", "final.gohtml")
	if err != nil {
		panic(err)
	}
	tpls = t
}

func main() {
	http.DefaultServeMux.HandleFunc("/", WithLogger(serveIndex))

	http.DefaultServeMux.HandleFunc("/redirect", WithLogger(handleRedirect))

	http.DefaultServeMux.HandleFunc("/moved-permanently", WithLogger(handleMoved))

	http.DefaultServeMux.HandleFunc("/final", WithLogger(serveResult))

	fmt.Println("Redirection service is listening on localhost:3000")
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
	tpls.ExecuteTemplate(w, "index.gohtml", nil)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		time.Sleep(time.Second * 1)
		http.Redirect(w, r, "/redirect", 303)
		return
	}
	// Temporary redirection
	time.Sleep(time.Second * 1)
	http.Redirect(w, r, "/moved-permanently", 307)
}

func handleMoved(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 1)
	http.Redirect(w, r, "/final", 301)
}

func serveResult(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 1)
	tpls.ExecuteTemplate(w, "final.gohtml", nil)
}
