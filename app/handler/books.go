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

	book, err := findOrFail(params["id"])

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(*book)
}

// CreateBook creates a new book to the local store
func CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	book := Book{ID: "test"}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(books)
}

// UpdateBook updates a book in the local store
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	newSetOfBooks := books[:0]

	if _, err := findOrFail(params["id"]); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

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

	if _, err := findOrFail(params["id"]); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	for _, book := range books {
		if book.ID != params["id"] {
			newSetOfBooks = append(newSetOfBooks, book)
		}
	}

	books = newSetOfBooks
	json.NewEncoder(w).Encode(books)
}

func findOrFail(id string) (*Book, error) {
	var book *Book
	for _, b := range books {
		if b.ID == id {
			book = &b
			return &b, nil
		}
	}

	if book == nil {
		return nil, fmt.Errorf("Book Not Found")
	}

	return &Book{}, nil
}
