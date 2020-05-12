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

// MockBooks creates mock data
func MockBooks() {
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
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook fetches single book from the local store
func GetBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	book := Book{ID: "test"}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		fmt.Println(err)

		panic(err)
	}

	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// UpdateBook updates a book in the local store
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	newSetOfBooks := books[:0]

	for _, book := range books {
		if book.ID == params["id"] {
			book := Book{ID: params["id"]}

			if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
				panic(err)
			}

			newSetOfBooks = append(newSetOfBooks, book)
		} else {
			newSetOfBooks = append(newSetOfBooks, book)
		}
	}

	books = newSetOfBooks
	json.NewEncoder(w).Encode(books)
}

// DeleteBook removes a book from the local store
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	newSetOfBooks := books[:0]

	for _, book := range books {
		if book.ID != params["id"] {
			newSetOfBooks = append(newSetOfBooks, book)
		}
	}

	books = newSetOfBooks
	json.NewEncoder(w).Encode(books)
}
