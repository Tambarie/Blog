package models

import (
	"net/http"
)


func CreateArticle(rw http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	author := r.PostFormValue("author")
	title := r.PostFormValue("title")
	article := r.PostFormValue("article")

	posts := &Post{
		Author:  author,
		Title:   title,
		Article: article,
	}

	// storing the data
	store(posts)

}
