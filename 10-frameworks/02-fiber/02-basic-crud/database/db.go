package database

import models "hanutsingh.in/learngo/frameworks/fiber/02/database/models"

var UserStore = models.UserList{
	"superadmin": models.User{
		FirstName: "Super",
		LastName:  "Admin",
		Password:  "qweasd@123",
	},
}
