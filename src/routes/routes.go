package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CustomRoute struct
type CustomRoute struct {
	url     string
	handler func(http.ResponseWriter, *http.Request)
	method  string
}

//CustomRoutes structs
type CustomRoutes []CustomRoute

func getRoutes() []CustomRoute {

	routes := []CustomRoutes{
		GetPostRoutes(),
	}

	var ApplicationRoutes []CustomRoute
	for _, r := range routes {
		ApplicationRoutes = append(ApplicationRoutes, r...)
	}
	return ApplicationRoutes

}

//New custom router
func New() *mux.Router {

	routes := getRoutes()

	// Set up routes
	router := mux.NewRouter()
	for _, route := range routes {
		router.HandleFunc(route.url, route.handler).Methods(route.method)
	}

	return router
}
