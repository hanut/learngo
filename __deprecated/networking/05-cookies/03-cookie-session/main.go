package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

const SessionDuration = 60 * 60 * 24

var tpls *template.Template

func init() {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}
	tpls = t
}

func main() {
	http.DefaultServeMux.HandleFunc("/", WithSession(serveIndex))

	http.DefaultServeMux.HandleFunc("/session/start", WithSession(handleStartSession))

	http.DefaultServeMux.HandleFunc("/session/end", WithSession(handleEndSession))

	http.DefaultServeMux.HandleFunc("/favicon.ico", http.NotFound)

	fmt.Println("App listening on localhost:3000...")
	http.ListenAndServe("localhost:3000", http.DefaultServeMux)

}

// WithSession returns a new handler function that wraps the handler passed in as an argument
// with a session variable
func WithSession(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var session MySession
		err := session.FromCookie(r)
		if err == nil {
			ctxWithSession := context.WithValue(r.Context(), "session", session)
			handlerFunc(w, r.WithContext(ctxWithSession))
			return
		}
		handlerFunc(w, r)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	session, ok := r.Context().Value("session").(MySession)
	if !ok {
		tpls.ExecuteTemplate(w, "index.gohtml", nil)
		return
	}
	tpls.ExecuteTemplate(w, "index.gohtml", struct{ Session MySession }{Session: session})
}

func handleStartSession(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("session").(MySession)
	if ok {
		fmt.Println("Session already exists in...")
		http.Redirect(w, r, "/?error=sessionexists", http.StatusTemporaryRedirect)
		return
	}
	ms := MySession{
		Id:        "10001",
		User:      "test.user@example.com",
		StartedAt: time.Now(),
	}
	v, err := ms.String()
	if err != nil {
		fmt.Println("ERROR:", err)
		http.Redirect(w, r, "/?error=internal", http.StatusTemporaryRedirect)
		return
	}
	sc := http.Cookie{
		Name:     "session",
		Value:    v,
		Path:     "/",
		MaxAge:   SessionDuration,
		HttpOnly: true,
	}
	http.SetCookie(w, &sc)
	if err != nil {
		fmt.Println("ERROR:", err)
		http.Redirect(w, r, "/?error=internal", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func handleEndSession(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("session").(MySession)
	if !ok {
		http.Redirect(w, r, "/?error=nosession", http.StatusTemporaryRedirect)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

type MySession struct {
	Id        string
	User      string
	StartedAt time.Time
}

func (s MySession) String() (string, error) {
	bd, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	cv := base64.StdEncoding.Strict().EncodeToString(bd)
	return string(cv), nil
}

func (s *MySession) FromCookie(r *http.Request) error {
	sc, err := r.Cookie("session")
	if err != nil {
		return err
	}
	cv, err := base64.StdEncoding.Strict().DecodeString(sc.Value)
	if err != nil {
		return err
	}
	err = json.Unmarshal(cv, s)
	if err != nil {
		return err
	}
	return nil
}
