package controllers

import (
	"fmt"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Book created")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of books")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get a single book")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update a book")
}
