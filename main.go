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

func (s *PostStore) Create(title, body string) Post {
	s.mu.Lock()
	defer s.mu.Unlock()
	p := Post{
		ID:    s.nextID,
		Title: title,
		Body:  body,
	}

	s.posts[p.ID] = p
	return p
}

func main() {
	store := newPostStore()

	post := store.Create("Hello World", "My first blog post")
	post2 := store.Create("Go is Fun", "I am learning Go")

	fmt.Println(post)
	fmt.Println(post2)
}
