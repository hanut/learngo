package database

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

// Setup global variables used in this module
var (
	// Session holds a reference to the mgo session
	Session *mgo.Session

	// Conn stores the mongo db connection string
	Conn *mgo.DialInfo
)

const (
	// DefaultMongoUri is the default uri that is used incase not sent in via the environment
	// variable MONGO_URI
	DefaultMongoUri = "mongodb://localhost:27017/crudapp"
)

// Connect connects to mongodb
func Init() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = DefaultMongoUri
	}

	mongo, err := mgo.ParseURL(uri)

	s, err := mgo.Dial(uri)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}

	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)

	Session = s
	Conn = mongo

}

func Close() {
	Session.Close()
}
