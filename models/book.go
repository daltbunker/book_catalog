package models

import (
	"gorm.io/gorm"
	"github.com/daltbunker/books/config"
	"encoding/csv"
	"os"
	"log"
	// "fmt"
)

var db * gorm.DB

type Book struct {
	gorm.Model
  BookId string `json:"bookId"`
	// GoodreadsBookId string `csv:"goodreads_book_id"`
	// BestBookId string `csv:"best_book_id"`
	// WorkId string `csv:"work_id"`
	// BooksCount string `csv:"books_count"`
	Isbn string `json:"isbn"`
	// Isbn13 string `csv:"isbn13"`
	Authors string `json:"authors"`
	OriginalPublicationYear string `json:"originalPublicationYear"`
	OriginalTitle string `json:"originalTitle"`
	// Title string `csv:"title"`
	// LanguageCode string `csv:"language_code"`
	AverageRating string `json:"averageRating"`
	RatingCount string `json:"ratingCount"`
	// WorkRatingsCount string `csv:"work_ratings_count"`
	// WorkTextReviewsCount string `csc:"work_text_reviews_count"`
	Ratings1 string `json:"ratings1"`
	Ratings2 string `json:"ratings2"`
	Ratings3 string `json:"ratings3"`
	Ratings4 string `json:"ratings4"`
	Ratings5 string `json:"ratings5"`
	ImageUrl string `json:"imageUrl"`
	// SmallImageUrl string `csv:"small_image_url"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int) *Book {
	var book Book
	db.First(&book, "ID = ?", id)
	return &book
}

func DeleteBook(id int) {
	var book Book
	db.Delete(&book, "ID = ?", id)
}

func UploadBooks(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	
	books := make([]Book, len(records) - 1)
	// range + 1 to avoid [first_name, last_name, username]
	for i := 1; i < len(records); i++ {
		books[i - 1] = Book{
			BookId: records[i][0],
			// GoodreadsBookId: records[i][1],
			// BestBookId: records[i][2],
			// WorkId: records[i][3],
			// BooksCount: records[i][4],
			Isbn: records[i][5],
			// Isbn13: records[i][6],
			Authors: records[i][7],
			OriginalPublicationYear: records[i][8],
			OriginalTitle: records[i][9],
			// Title: records[i][10],
			// LanguageCode: records[i][11],
			AverageRating: records[i][12],
			RatingCount: records[i][13],
			// WorkRatingsCount: records[i][14],
			// WorkTextReviewsCount: records[i][15],
			Ratings1: records[i][16],
			Ratings2: records[i][17],
			Ratings3: records[i][18],
			Ratings4: records[i][19],
			Ratings5: records[i][20],
			ImageUrl: records[i][21],
			// SmallImageUrl: records[i][22],
		}
	}

	db.Create(&books)
}