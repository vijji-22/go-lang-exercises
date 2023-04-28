package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
   "github.com/gorilla/mux"
)

type Book struct {
    Id string  "json:'id'"
    Name string  "json:'name'"
    Isbn string "json:'isbn'"
    Discount  string "json:'discount'"
    Typeofthebook string "json:'typeofthebook'"
    Author *Author  "json:'author'"
    Price *Price  "json:'price'"
}

type Author struct {
    Id string "json:'id'"
    Name  string "json:'name'"
    Noofbookswritten string  "json:'noofbookwritten'"
    Place  string "json:'place'"
    Language  string "json:'language'"
}

type Price struct{
    Onlineprice  string  "json:'onlineprice'"
    Offlineprice string "json:'offlineprice'"
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r) 

    for _, item := range books {
        if item.Id == params["ID"] {
            json.NewEncoder(w).Encode(item)
            return
        }

    }
    json.NewEncoder(w).Encode("Item not found")
}

func createBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var book Book
    json.NewDecoder(r.Body).Decode(&book)
    books = append(books, book)
    json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    for index, item := range books {
        if item.Id == params["ID"] {
            books = append(books[:index], books[index+1:]...)
            var book Book
            json.NewDecoder(r.Body).Decode(&book)
            books = append(books, book)
            json.NewEncoder(w).Encode(book)
            return
        }
    }
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    for index, item := range books {
        if item.Id == params["ID"] {
            books = append(books[:index], books[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(books)

}

var books []Book

func main() {
    r := mux.NewRouter()

    books = append(books, Book {Id: "1", Name: "Iskon", Isbn: "10", Discount: "10", Typeofthebook: "Bhagavathgeetha", Author: &Author{Id: "7", Name:"Valmki", Noofbookswritten: "10", Place: "Hyderabad", Language: "Telugu"}, Price: &Price{ Onlineprice: "250", Offlineprice:"300"}})
    books = append(books, Book {Id: "3", Name: "Twinkle", Isbn: "20", Discount: "40", Typeofthebook: "Geetha", Author: &Author{Id: "4", Name:"Tagore", Noofbookswritten: "20",Place: "Vizag",Language: "English"}, Price: &Price{ Onlineprice: "470",Offlineprice: "500"}})
    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/book/{ID}", getBook).Methods("GET")
    r.HandleFunc("/books", createBook).Methods("POST")
    r.HandleFunc("/books/{ID}", updateBook).Methods("PUT")
    r.HandleFunc("/books/{ID}", deleteBook).Methods("DELETE")

    fmt.Println("Starting server on 8080")
    log.Fatal(http.ListenAndServe(":8080", r))                     

}
