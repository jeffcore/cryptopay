package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	MongoDBUrl = "mongodb://hdwallet-mongo:27017"
)

// Connect connects to mongodb
func Connect() {
	uri := MongoDBUrl

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}


//
//
// func getSession() *mgo.Session {
//     // Connect to our local mongo
//     s, err := mgo.Dial("mongodb://hdwallet-mongo:27017")
//
//     // Check if connection error, is mongo running?
//     if err != nil {
//         panic(err)
//     }
//     defer s.Close()
//     return s.Clone()
// }
