package main

import "fmt"

type Post struct{
	ID int
	Title string
	Content string
	Author string
}

func main() {
post := Post{
	ID: 1,
	Title: "Learning Go",
	Content: "Go is a great backend language",
	Author: "Kevin",
}

fmt.Println(post)
}
