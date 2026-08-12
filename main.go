package main

import (
	"log"
	"net/http"

	"github.com/sravn25/blogo/internal/posts"
)

func main() {
	store := posts.NewStore()
	handler := posts.NewHandler(store)

	http.HandleFunc("GET /posts", handler.List)
	http.HandleFunc("POST /posts", handler.Create)
	http.HandleFunc("GET /posts/{id}", handler.Get)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
