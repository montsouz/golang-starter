package routes

import "../handlers"

//GetPostRoutes to get some post routes
func GetPostRoutes() []CustomRoute {
	var postRoutes = []CustomRoute{
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

	return postRoutes

}
