package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Structure
type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Pages int    `json:"pages"`
}

var books []Book

func mockBooks() {
	books = append(books, Book{
		ID:    "sw",
		Title: "Wait for me Angela",
		Pages: 128,
	})

	books = append(books, Book{
		ID:    "wf",
		Title: "Across the bridge",
		Pages: 324,
	})
}

// GetBooks fetches all the books from the local store
func GetBooks(w http.ResponseWriter, r *http.Request) {
	mockBooks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook fetches single book from the local store
func GetBook(w http.ResponseWriter, r *http.Request) {
	mockBooks()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// CreateBook creates a new book to the local store
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book := Book{ID: "test"}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		fmt.Println(err)

		panic(err)
	}

	books = append(books, book)
}

// UpdateBook updates a book in the local store
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)

}

// DeleteBook removes a book from the local store
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newSetOfBook := []Book{}

}

func getBookOr404(id string, w http.ResponseWriter, r *http.Request) *Book {
	var book *Book
	fmt.Println(book)

	for i, v := range books {
		fmt.Println(i, v, book)
		if book != nil {
			return book
		}

		if id == v.ID {
			println("Match", i)
			book = &v
		}
	}

	return book
}
