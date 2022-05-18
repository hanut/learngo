package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("ServeMux in Golang")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Sending back datetime")
		fmt.Println("Http Method Used:", r.Method)
		w.Write([]byte(time.Now().String()))
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Sending back users")
		fmt.Println("Http Method Used:", r.Method)
		w.Write([]byte("Alice, Bob, Charlie,David,Englebert,Henry"))
	})

	mux.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Sending back users")
		fmt.Println("Http Method Used:", r.Method)
		w.Write([]byte("Tommy, Moti, Sheena, Billu, Bhura, Kalu"))
	})

	mux.HandleFunc("/congrats", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "OOGA BOOGA WOOGA", 408)
			return
		}
		myPage := `
			<html>
				<head>
					<title>Congratulations !!!</title>
				</head>
				<body>
					<h1>Congratulations</h1>
					<p> You did it !!! </p>
				</body>
			</html>
		`
		w.Write([]byte(myPage))
	})

	fmt.Println("Our ServeMux is listening on localhost:8081")
	http.ListenAndServe("localhost:8081", mux)

}
