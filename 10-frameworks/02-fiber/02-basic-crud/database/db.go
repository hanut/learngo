package database

import models "hanutsingh.in/learngo/frameworks/fiber/02/database/models"

var UserStore = models.UserList{
	"superadmin": models.User{
		FirstName: "Super",
		LastName:  "Admin",
		Password:  "qweasd@123",
		Role:      "admin",
	},
	"hanut1": models.User{
		FirstName: "Hanut",
		LastName:  "Singh Gusain",
		Password:  "qweasd@123",
		Age:       33,
		Address:   "12-A, Nemi Road, Dalan Wala, Dehradun, UK, India",
		Role:      "manager",
	},
	"gurleen101010": models.User{
		FirstName: "Gurleen",
		LastName:  "Kaur",
		Password:  "qweasd@123",
		Age:       28,
		Address:   "35-Helene Crescent, Waterloo, Ontario, Canada",
		Role:      "employee",
	},
}
