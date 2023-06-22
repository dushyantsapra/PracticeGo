package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Pages int `json:"pages"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book := Book {
		Title: "Book Title",
		Author: "Book Author",
		Pages: 200,
	}

	json.NewEncoder(w).Encode(book)
}

func Hello(rw http.ResponseWriter, r *http.Request) {
	//rw.Header().Set("Content-Type", "test/html")
	rw.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/book", GetBook)

	log.Fatal(http.ListenAndServe(":5100", nil))
}
