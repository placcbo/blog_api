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
	Posts  map[int]Post
	nextID int
}

func newPostStore() *PostStore {

	return &PostStore{
		Posts:  map[int]Post{},
		nextID: 1,
	}
}

func (s *PostStore) Create(title, body string) Post {
	p := Post{
		ID:    s.nextID,
		Title: title,
		Body:  body,
	}
	s.nextID++
	s.Posts[p.ID] = p
	return p
}
func (s *PostStore) getAll() []Post {
	postsSlice := []Post{}

	for _, post := range s.Posts {
		postsSlice = append(postsSlice, post)
	}
	return postsSlice
}

func (s *PostStore) getPost(id int) (Post, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.Posts[id]
	return p, ok
}

func (s *PostStore) Update(id int, title, body string) (Post, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.Posts[id]
	if !ok {
		return Post{}, false
	}
	p.Title = title
	p.Body = body
	s.Posts[p.ID] = p
	return p, true

}

func main() {
	store := newPostStore()

	store.Create("golang series", "blog API")
	store.Create("golang series", "Todo API")

	updatePost, ok := store.Update(1, "books", "homo deus")
	fmt.Println(updatePost)
	fmt.Println(ok)
	fmt.Println(store.getAll())

}
