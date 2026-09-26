package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"sync"

	"github.com/go-chi/chi"
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

func NewPostStore() *PostStore {
	return &PostStore{
		posts:  map[int]Post{},
		nextID: 1,
	}
}

func (s *PostStore) CreatePost(title, body string) Post {
	s.mu.Lock()
	defer s.mu.Unlock()
	p := Post{
		ID:    s.nextID,
		Title: title,
		Body:  body,
	}
	s.posts[p.ID] = p
	s.nextID++
	return p

}

// get all posts

func (s *PostStore) GetAll() []Post {
	s.mu.Lock()
	defer s.mu.Unlock()
	p := []Post{}
	for _, post := range s.posts {
		p = append(p, post)
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].ID < p[j].ID
	})
	return p
}

// get single pos

func (s *PostStore) GetPost(id int) (Post, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.posts[id]
	if !ok {
		return Post{}, false
	}
	return p, true
}

// update post

func (s *PostStore) UpdatePost(id int, title, body string) (Post, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.posts[id]
	if !ok {
		return Post{}, false
	}
	p.Title = title
	p.Body = body

	s.posts[p.ID] = p
	return p, true

}

// delete

func (s *PostStore) DeletePost(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.posts[id]
	if !ok {
		return false
	}

	delete(s.posts, p.ID)
	return true
}

func main() {
	store := NewPostStore()

	r := chi.NewRouter()

	// create post

	r.Post("/posts", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid json from client", http.StatusNotFound)
			return
		}
		p := store.CreatePost(input.Title, input.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)
	})

	// get all posts

	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.GetAll())
	})

	//get post

	r.Get("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		p, ok := store.GetPost(idInt)
		if !ok {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	})

	// update post

	r.Put("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		var input struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		p, ok := store.UpdatePost(idNum, input.Title, input.Body)
		if !ok {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)
	})

	// delete

	r.Delete("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		p := store.DeletePost(idNum)
		json.NewEncoder(w).Encode(p)
	})

	http.ListenAndServe(":8080", r)

}
