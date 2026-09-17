package main

import (
	"fmt"
	"sync"
)

type Post struct {
	ID    int
	Title string
	Body  string
}

type PostStore struct {
	mu     sync.Mutex
	posts  map[int]Post
	nextID int
}

func main() {
	store := Post{}
	fmt.Println(store)
}
