package handlers

import (
	"net/http"
	"time"

	db "../db"
	models "../models"
	utils "../utils"
)

//CreatePost to create some post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{}
	utils.ReadBodyAndUnmarshalData(r, post)
	post.CreatedAt = time.Now().UTC()

	// Insert new post
	db.Insert(post, "posts", w)

}

//ReadPosts to read some posts
func ReadPosts(w http.ResponseWriter, r *http.Request) {
	sorting := "-created_at"
	posts := &[]models.Post{}
	db.Get(posts, "posts", sorting, w)
}
