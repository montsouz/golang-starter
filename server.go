package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	//Framework

	routes "./routes"
)

func main() {
	r := routes.New()
	http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
	log.Println("Listening on port 8080...")
}
