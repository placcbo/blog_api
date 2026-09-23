package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
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

	store.Create("golang projects todo", "todo api")
	store.Create("golang projects todo", "blogpost api")

	store.getAll()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.getAll())
	})

	http.ListenAndServe(":8080", mux)

}
