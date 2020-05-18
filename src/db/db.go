package db

import (
	"net/http"

	context "../context"
	utils "../utils"
	"gopkg.in/mgo.v2"
)

//GetModel of a db
func GetModel(model string) *mgo.Collection {
	session := context.GetMongoSession()
	return session.DB("app").C(model)
}

//Insert data in a db model
func Insert(data interface{}, modelName string, w http.ResponseWriter) {
	collection := GetModel(modelName)

	// Insert new post
	if err := collection.Insert(data); err != nil {
		utils.ResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, data)

}

//Get data from a db model
func Get(data interface{}, modelName string, sort string, w http.ResponseWriter) {
	collection := GetModel(modelName)

	if err := collection.Find(nil).Sort(sort).All(data); err != nil {
		utils.ResponseError(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.ResponseJSON(w, data)
	}
}
