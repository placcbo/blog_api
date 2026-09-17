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

func newPostStore() *PostStore {
	return &PostStore{
		posts:  map[int]Post{},
		nextID: 1,
	}
}


func main(){
	store := newPostStore()

	fmt.Println(store)
}
