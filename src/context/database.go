package context

import (
	"fmt"
	"log"
	"os"
	"sync"

	mgo "gopkg.in/mgo.v2"
)

var mongoSession *mgo.Session
var onceDatabase sync.Once

// GetMongoSession returns a copy of MongoDB session.
func GetMongoSession() *mgo.Session {
	onceDatabase.Do(func() {
		var err error

		mongoSession, err = getNewSession()
		if err != nil {
			errorMsg := fmt.Sprintf("Error on start database: %s", err.Error())
			panic(errorMsg)
		}
	})

	return mongoSession.Copy()
}

//Init my db
func getNewSession() (*mgo.Session, error) {
	// Connect to mongo
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("mongo err")
		os.Exit(1)
	}
	return session, err
}
