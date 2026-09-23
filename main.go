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
		Title: title,
		Body:  body,
		ID:    s.nextID,
	}

	s.posts[p.ID] = p
	s.nextID++
	return p
}

func (s *PostStore) All() []Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	var out []Post
	for _, p := range s.posts {
		out = append(out, p)
	}
	return out
}

func (s *PostStore) Get(id int) (Post, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.posts[id]
	return p, ok
}

func main() {
	store := newPostStore()

	store.Create("Hello world", "this is my first blog post")
	store.Create("why im learning go", "to get a good job someday!")

	post, ok := store.Get(1)
	fmt.Println(post)
	fmt.Println(ok)

}
