package main

import (
	"fmt"
	"hanut/learngo/gin/crudapp/database"
	"hanut/learngo/gin/crudapp/libs/auth"
	"hanut/learngo/gin/crudapp/middleware"
	"hanut/learngo/gin/crudapp/routes"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateUserForm struct {
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required"`
	Fname    string    `json:"fname" binding:"required,alpha"`
	Lname    string    `json:"lname" binding:"required,alpha"`
	Dob      time.Time `json:"dob" binding:"required"`
}

type UpdateUserForm struct {
	Fname string    `json:"fname" binding:"required,alpha"`
	Lname string    `json:"lname" binding:"required,alpha"`
	Dob   time.Time `json:"dob" binding:"required"`
}

type User struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email  string             `json:"email" bson:"email,omitempty"`
	Secret string             `json:"-" bson:"secret,omitempty"`
	Fname  string             `json:"fname" bson:"fname,omitempty"`
	Lname  string             `json:"lname" bson:"lname,omitempty"`
	Dob    primitive.DateTime `json:"dob" bson:"dob,omitempty"`
	Status bool               `json:"status" bson:"status,omitempty"`
}

func main() {
	// Setup the database
	err := database.InitMongo()
	if err != nil {
		log.Fatalln(err)
	}
	// Defer the close of db connection to after main ends
	defer database.CloseMongo()

	// Set mode to release for prod
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine
	r := gin.New()

	// Setup middleware
	r.Use(gin.Logger())

	// Custom example middleware
	r.Use(middleware.RequestIdMiddleware())
	r.Use(middleware.Logger())

	// Security Headers middleware
	r.Use(helmet.Default())
	r.Use(gin.Recovery())

	// Authenticated routes protected by a simple router middleware
	r.GET("/manage/account", routes.ManageAccount)
	r.GET("/manage/profile", routes.ManageMyProfile)

	// Endpoint for creating a new user
	r.POST("/users", func(ctx *gin.Context) {
		var uf CreateUserForm
		if err := ctx.ShouldBindJSON(&uf); err != nil {
			handleError(ctx, http.StatusBadRequest, err)
			return
		}
		s := auth.GenerateSecret(uf.Password)
		u := User{
			Email:  uf.Email,
			Secret: s,
			Fname:  uf.Fname,
			Lname:  uf.Lname,
			Dob:    primitive.NewDateTimeFromTime(uf.Dob),
			Status: false,
		}

		// Insert into the database
		r, e := database.ColUser.InsertOne(ctx.Request.Context(), u)
		if e != nil {
			handleError(ctx, 500, errors.Wrap(e, "Error adding new user"))
			return
		}
		// Retrieve the new user from the database and decode it into the user
		// Handle any errors if they occur
		e = database.ColUser.FindOne(ctx.Request.Context(), bson.M{"_id": r.InsertedID}).Decode(&u)
		if e != nil {
			handleError(ctx, 500, errors.Wrap(e, "Error finding newly added user"))
			return
		}
		ctx.JSON(201, u)
	})

	// Endpoint for getting a list of all the users
	r.GET("/users", func(ctx *gin.Context) {
		cs, e := database.ColUser.Find(ctx.Request.Context(), bson.M{})
		if e != nil {
			handleError(ctx, http.StatusInternalServerError, errors.Wrap(e, "Error creating cursor for user list"))
			return
		}
		var results []User
		e = cs.All(ctx.Request.Context(), &results)
		if e != nil {
			handleError(ctx, http.StatusInternalServerError, errors.Wrap(e, "Error getting all records from user list cursor"))
			return
		}
		ctx.JSON(200, results)
	})

	// // Endpoint for getting a specific user by id
	r.GET("/users/:userid", func(ctx *gin.Context) {
		// uihx is the user id as a hexadecimal string
		// as sent by the client
		uihx := ctx.Param("userid")

		// objid is the uihx converted to a valid mongo
		// objectid
		objid, e := primitive.ObjectIDFromHex(uihx)
		if e != nil {
			handleError(ctx, http.StatusBadRequest, errors.Wrap(e, "Invalid object id"))
			return
		}
		var fu User
		e = database.ColUser.FindOne(ctx.Request.Context(), bson.M{"_id": objid}).Decode(&fu)
		if e != nil {
			handleError(ctx, http.StatusNotFound, errors.Wrap(e, "No user found with that id"))
			return
		}

		ctx.JSON(200, fu)
	})

	r.PUT("/users/:userid", func(ctx *gin.Context) {
		// uihx is the user id as a hexadecimal string
		// as sent by the client
		uihx := ctx.Param("userid")

		// objid is the uihx converted to a valid mongo
		// objectid
		objid, e := primitive.ObjectIDFromHex(uihx)
		if e != nil {
			handleError(ctx, http.StatusBadRequest, errors.Wrap(e, "Invalid object id"))
			return
		}
		var uf UpdateUserForm
		if err := ctx.BindJSON(&uf); err != nil {
			handleError(ctx, 400, err)
			return
		}
		var o options.FindOneAndUpdateOptions
		o.SetReturnDocument(options.After)

		r := database.ColUser.FindOneAndUpdate(ctx.Request.Context(), bson.M{"_id": objid}, bson.M{
			"$set": bson.M{
				"fname": uf.Fname,
				"lname": uf.Lname,
				"dob":   uf.Dob,
			},
		}, &o)
		if r.Err() != nil {
			handleError(ctx, 403, r.Err())
			return
		}
		var user User
		r.Decode(&user)
		ctx.JSON(200, user)
	})

	r.PATCH("/users/:userid/status", func(ctx *gin.Context) {
		// uihx is the user id as a hexadecimal string
		// as sent by the client
		uihx := ctx.Param("userid")

		// objid is the uihx converted to a valid mongo
		// objectid
		objid, e := primitive.ObjectIDFromHex(uihx)
		if e != nil {
			handleError(ctx, http.StatusBadRequest, errors.Wrap(e, "Invalid object id"))
			return
		}
		var o options.FindOneAndUpdateOptions
		o.SetReturnDocument(options.After)

		r := database.ColUser.FindOneAndUpdate(ctx.Request.Context(), bson.M{"_id": objid}, []bson.M{
			{
				"$set": bson.M{
					"status": bson.M{"$not": "$status"},
				},
			},
		}, &o)
		if r.Err() != nil {
			handleError(ctx, 403, r.Err())
			return
		}
		var user User
		r.Decode(&user)
		ctx.JSON(200, user)
	})

	r.DELETE("/users/:userid", func(ctx *gin.Context) {
		uid := ctx.Param("userid")
		objid, e := primitive.ObjectIDFromHex(uid)

		if e != nil {
			handleError(ctx, http.StatusBadRequest, errors.Wrap(e, "Invalid object id"))
			return
		}

		r := database.ColUser.FindOneAndDelete(ctx.Request.Context(), bson.M{"_id": objid})
		e = r.Err()
		if e != nil {
			if errors.Is(e, mongo.ErrNoDocuments) {
				handleError(ctx, 404, errors.New("User not found"))
				return
			}
			handleError(ctx, 500, r.Err())
			return
		}
		var user User
		r.Decode(&user)
		ctx.JSON(200, user)
	})

	p, e := strconv.ParseInt(os.Getenv("CA_PORT"), 10, 16)
	if e == nil {
		p = 3000
	}
	r.Run(fmt.Sprintf(":%d", p))
}

func handleError(ctx *gin.Context, ecode int, e error) {
	ctx.JSON(ecode, gin.H{"error": e.Error()})
	log.Printf("[ERROR] (%d) %s \n", ecode, e.Error())
}
