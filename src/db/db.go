package db

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

//Init my db
func Init() {
	// Connect to mongo
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("mongo err")
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
}

//GetSession of a db
func GetSession() *mgo.Session {
	return session.Copy()
}

//GetModel of a db
func GetModel(model string) *mgo.Collection {
	return session.DB("app").C(model)
}
