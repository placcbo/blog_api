package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
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

// create post

func (s *PostStore) CreatePost(title, body string) Post {
	p := Post{
		Title: title,
		Body:  body,
		ID:    s.nextID,
	}
	s.posts[p.ID] = p
	s.nextID++
	return p
}

func (s *PostStore) GetAll() []Post {
	p := []Post{}

	for _, post := range s.posts {
		p = append(p, post)
	}
	return p
}

func (s *PostStore) GetPost(id int) (Post, bool) {
	post, ok := s.posts[id]
	if !ok {
		return Post{}, false
	}
	return post, true

}

func main() {
	store := newPostStore()
	fmt.Println(store.GetPost(2))

	r := chi.NewRouter()

	// get all posts
	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.GetAll())
	})

	// get post

	r.Get("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid not found", http.StatusNotFound)
			return
		}
		post, ok := store.GetPost(idNum)
		if !ok {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(post)
	})

	http.ListenAndServe(":8080", r)

}
