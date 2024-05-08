package api

import (
	"encoding/json"
	"github.com/eduardoraider/go-crud/internal/entity"
	"net/http"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook entity.Book
	json.NewDecoder(r.Body).Decode(&newBook)

	entity.AddBook(newBook)

	w.WriteHeader(http.StatusCreated)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook entity.Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	entity.UpdateBook(updatedBook)

	w.WriteHeader(http.StatusOK)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	var bookID struct {
		ID string `json:"id"`
	}
	json.NewDecoder(r.Body).Decode(&bookID)

	entity.DeleteBook(bookID.ID)

	w.WriteHeader(http.StatusOK)
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", getBooks)
	mux.HandleFunc("/books/create", createBook)
	mux.HandleFunc("/books/update", updateBook)
	mux.HandleFunc("/books/delete", deleteBook)
	return mux
}
