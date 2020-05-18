package db

import (
	context "../context"
	"gopkg.in/mgo.v2"
)

//GetModel of a db
func GetModel(model string) *mgo.Collection {
	session := context.GetMongoSession()
	return session.DB("app").C(model)
}
