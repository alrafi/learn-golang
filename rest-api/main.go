package main

import (
	// "fmt"
	"log"
	"encoding/json"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	ID			string	`json:"id"`
	Isbn		string	`json:"isbn"`
	Title		string	`json:"title"`
	Author	*Author	`json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname 	string `json:"lastname"`
}

// init books
var books []Book

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) // get params

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// add a book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// update a book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init router
	r := mux.NewRouter()

	// mock data
	books = append(books, Book{ID: "1", Isbn: "34590127", Title: "Sky Full of Stars", Author: &Author{Firstname: "Chris",Lastname: "Martin"}})
	books = append(books, Book{ID: "2", Isbn: "89890342", Title: "After Sunset Tomorrow", Author: &Author{Firstname: "Evan", Lastname: "Smith"}})

	// route handlers (endpoints)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}