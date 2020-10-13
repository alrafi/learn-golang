package main

import (
	// "fmt"
	"log"
	// "encoding/json"
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

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get a book
func getBook(w http.ResponseWriter, r *http.Request) {

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

	// route handlers (endpoints)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}