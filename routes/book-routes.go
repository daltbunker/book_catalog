package routes

import (
	"github.com/daltbunker/books/controllers"
	"github.com/go-chi/chi/v5"
)

var RegisterBookStoreRoutes = func(router *chi.Mux) {
	router.Post("/books/", controllers.CreateBook)
	router.Get("/books/", controllers.GetBooks)
	router.Get("/books/{bookId}", controllers.GetBookById)
	router.Delete("/books/{bookId}", controllers.DeleteBook)
	router.Post("/books/upload/", controllers.UploadBooks)
}