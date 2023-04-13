package database

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	FirstName string `validate:"required,min=3,max=32"`
	LastName  string `validate:"required,min=3,max=32"`
	Password  string `validate:"required,min=8,max=16,ascii"`
	Age       uint8  `validate:"required,gte=0,lte=150"`
	Address   string `validate:"required,min=5,max=256"`
	Role      string `validate:"required,alpha,oneof='admin' 'manager' 'employee',min=5,max=12"`
}

type UserList map[string]User

// GetUser returns a specific user from the list of users
func (ul UserList) GetUser(id string) (User, error) {
	u, ok := ul[id]
	if !ok {
		return u, errors.New("User not found")
	}
	return u, nil
}

// Add user adds a new user into the list of users or throws
// and error if the user already exists
func (ul UserList) AddUser(u User) (string, error) {
	data := []byte(u.FirstName + u.LastName + u.Address + fmt.Sprintf("%d", u.Age))
	gen := uuid.NewSHA1(uuid.UUID{}, data)
	hash := gen.String()
	_, ok := ul[hash]
	if ok {
		return hash, errors.New("User already exists in the database")
	}

	ul[hash] = u
	return hash, nil
}

// RemoveUser removes a user by id from the list of users
func (ul UserList) RemoveUser(id string) error {
	_, ok := ul[id]
	if !ok {
		return errors.New("User not found in database")
	}
	delete(ul, id)
	return nil
}

// ReplaceUser searches for a user by id and replaces the values with
// the new user object passed in as an argument
func (ul UserList) ReplaceUser(id string, u User) error {
	_, ok := ul[id]
	if !ok {
		return errors.New("Error: User not found")
	}
	fmt.Println("Before Switch", ul[id])
	ul[id] = u
	fmt.Println("After Switch", ul[id])
	return nil
}
