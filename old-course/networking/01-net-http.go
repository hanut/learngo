package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Person struct {
	Name  string
	Age   uint8
	Email string
}

type People []Person

var plist People = People{
	Person{Name: "John", Age: 22, Email: "john.smith@example.com"},
	Person{Name: "Mary", Age: 29, Email: "mary.sue@example.com"},
	Person{Name: "Bob", Age: 30, Email: "bob@example.com"},
	Person{Name: "Alice", Age: 31, Email: "alice@example.com"},
	Person{Name: "Tim", Age: 70, Email: "tim.tiny@example.com"},
	Person{Name: "Gulshan", Age: 35, Email: "gul.shan.22@example.com"},
	Person{Name: "Kabir", Age: 42, Email: "kabir.singh@example.com"},
	Person{Name: "Vladimir", Age: 55, Email: "vlad.imp@example.com"},
	Person{Name: "Scrooge", Age: 28, Email: "scrooge.dollar@example.com"},
}

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

	func(res *http.ResponseWriter, req *http.Request) {
		js, err := json.Marshal(plist)

		if strings.Contains(req.URL.Path, "error") {
			(*res).WriteHeader(400)
			(*res).Write([]byte("Bad Request"))
			return
		}

		if err != nil {
			(*res).WriteHeader(500)
			(*res).Write([]byte("Internal Server Error"))
			return
		}

		(*res).WriteHeader(200)
		(*res).Write(js)
	}(&res, req)

	fmt.Printf("%s -> %s %s [%v]\n", req.RemoteAddr, req.Method, req.URL, time.Since(st))

}

func main() {
	app := myHttpServer{
		Host:    "localhost",
		Port:    3001,
		AppName: "People Microservice",
		Env:     "development",
	}

	fmt.Printf("%s (%s) listening at %s\n", app.AppName, app.Env, app.Address())
	http.ListenAndServe(app.Address(), &app)
}
