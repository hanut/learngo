package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var users []string = []string{"Alice", "Bob", "Charlie", "Dan"}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	jsd, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error:", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	_, err = w.Write(jsd)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}

func CreateUsers(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var data struct {
		Name string `json:"name"`
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Read Error:", err)
		http.Error(w, err.Error(), 500)
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Parse Error:", err)
		http.Error(w, err.Error(), 500)
		return
	}
	users = append(users, data.Name)
	w.Write([]byte("OK"))
}
