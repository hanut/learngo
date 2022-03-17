package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id     string
	Email  string
	Secret string
	Fname  string
	Lname  string
	Dob    time.Time
	Status bool
}

type UserList []User

// FindUserById looks for a user with the given id in the user list
// and returns it if found. If not, the f return value will be false
func (ul *UserList) FindUserById(uid string) (fu *User, f bool) {
	for _, u := range *ul {
		if u.Id == uid {
			f = true
			fu = &u
		}
	}
	return
}

func (ul *UserList) FindAndDeleteById(uid string) (err error) {
	fi := -1

	for i, u := range *ul {
		if u.Id == uid {
			fi = i
		}
	}

	if fi == -1 {
		err = errors.New("User not found")
		return
	}

	// Remove the user from the slice
	var tmp []User = append([]User{}, []User(*ul)[:fi]...) // Convert to slice so we can do stuff
	tmp = append(tmp, []User(*ul)[fi+1:]...)

	*ul = UserList(tmp) // Conver back to user list and assign to the original list
	return
}

func main() {
	// Setup the database
	// database.Init()
	// defer database.Close()

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	var ul UserList

	// Endpoint for getting a list of all the users
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, ul)
	})

	// Endpoint for getting a specific user by id
	r.GET("/users/:userid", func(ctx *gin.Context) {
		uid := ctx.Param("userid")
		println(uid, len(uid))
		if len(uid) != 24 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
			return
		}

		fu, f := ul.FindUserById(uid)
		if !f {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		ctx.JSON(200, *fu)
	})

	// Endpoint for getting a specific user by id
	r.POST("/users", func(ctx *gin.Context) {
		var u User
		if err := ctx.ShouldBindJSON(&u); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ul = append(ul, u)
		ctx.JSON(201, true)
	})

	r.PUT("/users", func(ctx *gin.Context) {
		var u User
		if err := ctx.ShouldBindJSON(&u); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fu, f := ul.FindUserById(u.Id)
		if !f {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
			return
		}

		// Update the user by replacing value of the pointer with new user
		// received from client
		*fu = u

		ctx.JSON(200, *fu)
	})

	r.DELETE("/users/:userId", func(ctx *gin.Context) {
		uid := ctx.Param("userId")

		if len(uid) != 24 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
			return
		}

		err := ul.FindAndDeleteById(uid)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, true)
	})

	r.Run("localhost:3000")
}
