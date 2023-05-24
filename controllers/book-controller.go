package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
  "io"
	"github.com/daltbunker/books/models"
	"github.com/daltbunker/books/utils"
	"github.com/go-chi/chi/v5"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _  := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")
	parsedBookId, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error parsing book id")
	}
	book := models.GetBookById(parsedBookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")
	parsedBookId, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error parsing book id")
	}
	models.DeleteBook(parsedBookId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func UploadBooks(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileSize := header.Size
	fmt.Printf("File size (bytes): %v\n", fileSize)

	tmpFile , err := os.CreateTemp("", "temp-books-*.csv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	models.UploadBooks(tmpFile.Name())
	w.WriteHeader(200)
}