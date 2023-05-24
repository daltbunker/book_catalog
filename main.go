package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/daltbunker/books/routes"
)
func main() {
	r := chi.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}