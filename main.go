package main

import (
	"encoding/json"
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

func (s *PostStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Posts[id]; !ok {
		return false
	}

	delete(s.Posts, id)
	return true
}

func main() {
	store := newPostStore()

	store.Create("golang projects todo", "todo api")
	store.Create("golang projects todo", "blogpost api")

	r := chi.NewRouter()

	//GET POSTS

	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.getAll())
	})

	//POST posts

	r.Post("/posts", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		p := store.Create(input.Title, input.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)

	})

	r.Get("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		idNum, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		post, ok := store.getPost(idNum)
		if !ok {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)

	})

	// Update

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

		post, ok := store.Update(idNum, input.Title, input.Body)

		if !ok {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	})

	r.Delete("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		idNum, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		if !store.Delete(idNum) {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe(":8080", r)

}
