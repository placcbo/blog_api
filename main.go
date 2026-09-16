package main

import "fmt"

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

	post, found:= getPostByID(posts, 10)
	if found == false{
		fmt.Println("Post not found")
	}
	fmt.Println("Found post:")
	fmt.Println("ID:", post.ID)
	fmt.Println("Title:", post.Title)
	fmt.Println("Author:", post.Author)

}
