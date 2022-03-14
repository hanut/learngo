package main

import (
	"fmt"
	"time"
)

type User struct {
	Id          int32
	Email       string
	DateOfBirth time.Time
	Social      SocialCon
}

type SocialCon struct {
	Facebook bool
	Google   bool
	Twitter  bool
	Linkedin bool
	Github   bool
}

func main() {
	users := [3]User{
		{
			Id:          1,
			Email:       "alice@example.com",
			DateOfBirth: time.Now(),
			Social:      SocialCon{Facebook: true, Google: true, Twitter: false, Linkedin: false, Github: false},
		},
		{
			Id:          2,
			Email:       "bob@example.com",
			DateOfBirth: time.Now(),
			Social:      SocialCon{Facebook: false, Google: false, Twitter: false, Linkedin: true, Github: true},
		},
		{
			Id:          3,
			Email:       "charlie@example.com",
			DateOfBirth: time.Now(),
			Social:      SocialCon{Facebook: false, Google: false, Twitter: true, Linkedin: false, Github: false},
		},
	}

	// Create a map from array
	id2email := map[int32]string{}

	for _, u := range users {
		id2email[u.Id] = u.Email
	}

	fmt.Println("Id > Email", id2email)

	charliesId := 3
	fmt.Println("Charlie's email id:", id2email[int32(charliesId)])

	// Create a map using make
	id2Dob := make(map[int32]time.Time)

	for _, u := range users {
		id2Dob[u.Id] = u.DateOfBirth
	}
	fmt.Println("Charlie's Birthday:", id2Dob[int32(charliesId)])
}
