package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tpls *template.Template

func init() {
	// Load the go template
	t, err := template.ParseFiles("index.gohtml", "result.gohtml")
	if err != nil {
		panic(err)
	}
	tpls = t
}

func main() {

	http.DefaultServeMux.HandleFunc("/", ServeUploadPage)

	http.DefaultServeMux.HandleFunc("/upload-file", HandleFileUpload)

	fmt.Println("Server starting on localhost:3000...")
	http.ListenAndServe("localhost:3000", http.DefaultServeMux)
}

func HandleFileUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorNotFound(w)
		return
	}
	f, fh, err := r.FormFile("myFile")
	if err != nil {
		ErrorBadRequest(w)
		return
	}
	d, err := ioutil.ReadAll(f)
	if err != nil {
		ErrorInternalServerError(w, err.Error())
		return
	}
	tvars := map[string]string{
		"content":  string(d),
		"filename": fh.Filename,
		"filetype": fh.Header.Get("Content-Type"),
	}
	fmt.Println("Tvars", tvars)
	err = tpls.ExecuteTemplate(w, "result.gohtml", tvars)
	if err != nil {
		ErrorInternalServerError(w, err.Error())
	}
}

func ServeUploadPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorNotFound(w)
		return
	}
	tpls.ExecuteTemplate(w, "index.gohtml", nil)
}

func ErrorNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

func ErrorBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request"))
}

func ErrorInternalServerError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(message))
}
