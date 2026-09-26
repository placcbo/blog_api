# Blog API

A RESTful Blog API built with Go.

This project implements a complete CRUD API for managing blog posts using Go's
standard HTTP package, the Chi router, JSON, and an in-memory data store.

The project was built to practice backend fundamentals including HTTP routing,
request/response handling, concurrency protection, error handling, middleware,
and RESTful API design.

## Features

- Create a blog post
- Get all blog posts
- Get a single blog post by ID
- Update a blog post
- Delete a blog post
- JSON request and response handling
- HTTP status codes
- Request validation and error handling
- Chi HTTP routing
- Request logging middleware
- Thread-safe in-memory storage
- Mutex-protected shared state
- Automatic incremental post IDs
- Deterministic ordering when retrieving posts

## Technologies

- Go
- `net/http`
- Chi router
- JSON
- `sync.Mutex`
- In-memory storage
- Thunder Client
- Git & GitHub

## Project Structure

```text
blog_api/
├── main.go
├── go.mod
├── go.sum
└── README.md
