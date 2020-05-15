package routes

import "../handlers"

//GetPostRoutes to get some post routes
func GetPostRoutes() []CustomRoute {
	var postReoutes = []CustomRoute{
		CustomRoute{
			url:     "/posts",
			handler: handlers.CreatePost,
			method:  "POST",
		},
		CustomRoute{
			url:     "/posts",
			handler: handlers.ReadPosts,
			method:  "GET",
		},
	}

	return postReoutes

}
