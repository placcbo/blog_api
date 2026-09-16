package main

import "fmt"

type Post struct {
	ID      int
	Title   string
	Content string
	Author  string
}

func main() {
	posts := []Post{
		{
			ID:      1,
			Title:   "Learn golang",
			Content: "Learn about the blog post API",
			Author:  "Kevin",
		},
		{
			ID:      2,
			Title:   "Building APIS",
			Content: "APIs allow applications to communicate",
			Author:  "Kevin",
		},
	}

	for _, post := range posts{
		fmt.Println(post.Title)
	}

	
}
