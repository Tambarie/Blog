package models

type Post struct {
	Author string
	Title string
	Article string
}

var Entries  []*Post

func store(post *Post)  {
	Entries = append(Entries, post)
}