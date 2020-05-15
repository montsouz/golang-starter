package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	db "../db"
	models "../models"
)

//CreatePost to create some post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	posts := db.GetModel("posts")
	// Read body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Read post
	post := &models.Post{}
	err = json.Unmarshal(data, post)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.CreatedAt = time.Now().UTC()

	// Insert new post
	if err := posts.Insert(post); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON(w, post)
}

//ReadPosts to read some posts
func ReadPosts(w http.ResponseWriter, r *http.Request) {
	posts := db.GetModel("posts")
	result := []models.Post{}
	if err := posts.Find(nil).Sort("-created_at").All(&result); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
	} else {
		responseJSON(w, result)
	}
}
