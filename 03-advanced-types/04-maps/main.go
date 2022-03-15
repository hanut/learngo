package main

import "fmt"

type User struct {
	Id     int
	Status bool
}

func main() {
	fmt.Println("Working with maps in Golang")

	// A simple map
	simpleMap := map[string]string{
		"user1@example.com": "someRandomPassword1",
		"user2@example.com": "someRandomPassword2",
		"user3@example.com": "someRandomPassword3",
		"user4@example.com": "someRandomPassword4",
	}

	simpleMap["user5@example.com"] = "AnotherRandomPassword5"

	for k, v := range simpleMap {
		fmt.Println("User", k, "has password", v)
	}

	// Direct access the value by key
	fmt.Println("Password for user2@example.com is", simpleMap["user2@example.com"])

	// Convert array to map
	users := [3]User{
		{Id: 1, Status: false},
		{Id: 2, Status: false},
		{Id: 3, Status: true},
	}

	// The id we want to find
	idToCheck := 3

	// Find using range
	for _, v := range users {
		if v.Id == idToCheck {

			fmt.Println("Is User with id 3 Online ?", v.Status)
			break
		}
	}

	// Simple boolean map that stores online/offline status of user by id
	statusMap := map[int]bool{}
	for _, v := range users {
		statusMap[v.Id] = v.Status
	}

	fmt.Println("Is User with id 3 Online ?", statusMap[idToCheck])

}
