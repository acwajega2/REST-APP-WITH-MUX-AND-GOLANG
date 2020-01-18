package main

import (
//	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"

)
///CREATING THE STRUCTS ---MODEL
// Book Struct (Model)
type Book struct{
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Author *Author `json: "author"`
	
}

//----Author Struct
type Author struct{
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`	

}

//----Intit Books var as a Sliec Book struct
var books []Book

func home(w http.ResponseWriter, r*http.Request){

}

//----Get all books function
func getBooks(w http.ResponseWriter, r*http.Request){
w.Header().Set("Content-Type","application/json")
json.NewEncoder(w).Encode(books)

}
//--------------------------------------------
//----Get single book function
func getBook(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r) //---get all params

	//loop through books and find with id
	for _, item :=range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return

		}
		
		
	}
	json.NewEncoder(w).Encode(&Book{})
	
	}
	//--------------------------------------------
//----Create single Book
func createBook(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000)) //mock ID
	books = append(books,book)
	json.NewEncoder(w).Encode(book)
}
//--------------------------------------------

//----Update Book
func updateBook(w http.ResponseWriter, r*http.Request){
	
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item :=range books {
		if item.ID == params["id"]{
			books = append(books[:index],books[index+1:]...)
			var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000)) //mock ID
	books = append(books,book)
	return

		}
		json.NewEncoder(w).Encode(books)
	}

}
//--------------------------------------------

//----Delete Book unction
func deleteBook(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item :=range books {
		if item.ID == params["id"]{
			books = append(books[:index],books[index+1:]...)
			break

		}
		json.NewEncoder(w).Encode(books)
	}
	

}
//--------------------------------------------


func main(){


	//Intitalising the router
	r := mux.NewRouter()

	//Mock Data -@todo implement database
	books = append(books,Book{ID:"6",Isbn:"5532323",Title : "Book Two",Author : &Author{Firstname:"Deris",Lastname:"Matovu"}})
	books = append(books,Book{ID:"1",Isbn:"32323",Title : "Book one",Author : &Author{Firstname:"Chris",Lastname:"Wajega"}})
	
	//Route Handlers /End points
	r.HandleFunc("/",home).Methods("GET")
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))


}