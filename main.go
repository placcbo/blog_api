package main

import (
	"encoding/json"
	"net/http"
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

func (s *PostStore) Create(title, body string) Post {
	p := Post{
		Title: title,
		Body:  body,
		ID:    s.nextID,
	}

	s.posts[p.ID] = p
	s.nextID++
	return p
}

func (s *PostStore) GetPost(id int) (Post, bool) {
	p, ok := s.posts[id]
	if !ok {
		return Post{}, false
	}
	return p, ok
}

func (s *PostStore) Getall() []Post {
	p := []Post{}
	for i := 1; i < s.nextID; i++ {
		p = append(p, s.posts[i])
	}

	return p
}

func main() {
	store := NewPostStore()

	store.Create("Building Better REST APIs", "A good REST API uses clear endpoints, meaningful HTTP methods, consistent responses, and proper status codes. These principles make APIs easier to understand and maintain.")
	store.Create("Understanding Go for Backend Development", "Go provides a simple and efficient environment for building backend applications. Its performance, concurrency features, and straightforward syntax make it useful for API development.")
	store.Create("Why Clean Code Matters", "Clean code is easier to read, test, debug, and maintain. Simple functions, meaningful names, and consistent structure can make a large application much easier to work with.")
	store.Create("Introduction to Database Design", "A well-designed database helps applications store and retrieve information efficiently. Tables, relationships, indexes, and constraints all contribute to reliable data management.")
	store.Create("Getting Started With JSON", "JSON is a popular format for exchanging structured data between clients and servers. Its simple syntax makes it easy for both humans and applications to understand.")
	store.Create("API Authentication Explained", "Authentication allows an API to determine who is making a request. Common approaches include sessions, API keys, and token-based authentication.")
	store.Create("The Importance of Input Validation", "Input validation protects applications from invalid and unexpected data. APIs should validate incoming values before processing or storing them.")
	store.Create("Designing Pagination for APIs", "Pagination prevents APIs from returning unnecessarily large responses. Clients can request smaller groups of records, improving performance and reducing ntwork usage.")
	store.Create("Caching and Web Performance", "Caching can reduce repeated work by temporarily storing frequently requested data. When implemented carefully, it can improve response times and reduce database load.")
	store.Create("Testing Backend Applications", "Automated tests help developers verify that application behavior remains correct as the codebase changes. Unit tests and integration tests can catch problems early.")
	store.Create("Error Handling in APIs", "Clear error handling helps API clients understand what went wrong. Consistent status codes and structured error responses make debugging and integration easier.")
	store.Create("Building Scalable Applications", "Scalable applications are designed to handle increasing traffic and data without unnecessary complexity. Good architecture, caching, database optimization, and monitoring all contribute to scalability.")
	store.Create("Introduction to HTTP Methods", "HTTP methods describe the intended action of a request. GET commonly retrieves data, POST creates resources, PUT updates resources, and DELETE removes resources.")
	store.Create("Writing Effective API Documentation", "Good API documentation explains available endpoints, request parameters, response formats, authentication requirements, and common errors. Clear documentation makes an API easier for developers to use.")
	store.Create("Monitoring Backend Services", "Monitoring provides visibility into application performance and reliability. Metrics, logs, and alerts can help developers identify problems and understand system behavior.")
	store.Create("Security Best Practices for APIs", "API security involves protecting authentication credentials, validating input, controlling access, limiting abuse, and avoiding unnecessary exposure of sensitive information.")
	store.Create("Working With HTTP Status Codes", "HTTP status codes communicate the result of a request. Using appropriate codes helps clients distinguish successful requests from validation errors, authentication failures, missing resources, and server problems.")
	store.Create("The Role of Middleware", "Middleware provides reusable processing between an incoming request and the final handler. Logging, authentication, request validation, and recovery are common middleware responsibilities.")
	store.Create("From Prototype to Production", "Moving an application into production requires more than writing code. Configuration, testing, security, monitoring, deployment, and reliable data management all become important parts of the process.")

	r := chi.NewRouter()

	// get all

	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.Getall())
	})

	// get post

	r.Get("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		NumID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid id", http.StatusNotFound)
			return
		}

		p, ok := store.GetPost(NumID)
		if !ok {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)

	})

	http.ListenAndServe(":8080", r)

}
