package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"gopkg.in/mgo.v2"

	//Framework
	db "./src/db"
	routes "./src/routes"
)

var posts *mgo.Collection

func main() {
	db.Init()
	r := routes.New()
	http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
	log.Println("Listening on port 8080...")
}
