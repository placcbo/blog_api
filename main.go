package main

import (
	"fmt"
	"net/http"
)

type Post struct {
	ID      int
	Title   string
	Content string
	Author  string
}

func getPostByID(posts []Post, id int) (Post, bool) {
	for _, post := range posts {
		if post.ID == id {
			return post, true
		}
	}

	return Post{}, false
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my Blog API")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All blog posts")
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
			Title:   "Building APIs",
			Content: "APIs allow applications to communicate",
			Author:  "Kevin",
		},
	}

	post, found := getPostByID(posts, 10)

	if found {
		fmt.Println("Found post:")
		fmt.Println("ID:", post.ID)
		fmt.Println("Title:", post.Title)
		fmt.Println("Author:", post.Author)
	} else {
		fmt.Println("Post not found")
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/posts", postHandler)

	http.ListenAndServe(":8080", nil)
}
