package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

var tpls *template.Template

func init() {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}
	tpls = t
}

func main() {

	http.DefaultServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s, err := r.Cookie("session")
		var session MySession
		if err == nil {
			session.FromJSON(s.Value)
			tpls.ExecuteTemplate(w, "index.gohtml", struct{ Session MySession }{Session: session})
			return
		}
		tpls.ExecuteTemplate(w, "index.gohtml", nil)
	})

	http.ListenAndServe("localhost:3000", http.DefaultServeMux)

}

type MySession struct {
	Id        string
	User      string
	StartedAt time.Time
}

func (s MySession) String() (string, error) {
	s, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return s, nil
}

func (s *MySession) FromJSON(data string) error {
	err := json.Unmarshal(s, data)
	if err != nil {
		return err
	}
	return nil
}
