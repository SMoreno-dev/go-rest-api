package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books var as a slice Book struct
var books []Book

func main() {
	//Init Router
	r := mux.NewRouter()

	// Mock data - @todo - implement DB
	books = append(books, Book{
		ID:    "1",
		Isbn:  "545323",
		Title: "First Book",
		Author: &Author{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Isbn:  "647423",
		Title: "Second Book",
		Author: &Author{
			Firstname: "Steve",
			Lastname:  "Smith",
		},
	})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

//Get books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get params
	params := mux.Vars(r)

	//Loop through books and find id
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode("Book not found.")
}

//Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Decode request body into var
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	//Mock Id - NOT safe
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)

	//Write a response
	json.NewEncoder(w).Encode(&book)

}

//Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Delete book
	for i, book := range books {
		if book.ID == params["id"] {
			//Delete Book
			books = append(books[:i], books[i+1:]...)

			//Create new Book
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)

			//Response
			json.NewEncoder(w).Encode(book)
			return
		}
	}

}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Delete book
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	//Response
	json.NewEncoder(w).Encode(&books)
}
