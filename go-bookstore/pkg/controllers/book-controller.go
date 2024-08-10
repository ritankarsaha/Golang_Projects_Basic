package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/ritankarsaha/go-bookstore/pkg/utils"
	"github.com/ritankarsaha/go-bookstore/pkg/models"
)

//controller for retrieving all the books from the database
func GetBook(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, "Unable to retrieve books", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


// GetBookById retrieves a single book by ID from the database and returns it as JSON
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseUint(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := models.GetBookById(uint(ID))
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


// CreateBook creates a new book in the database
func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
 utils.ParseBody(r, book)
	createdBook, err := book.CreateBook()
	if err != nil {
		http.Error(w, "Error creating book", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(createdBook)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) 
	w.Write(res)
}


//delete book controller deletes a  book from the database
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseUint(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if err := models.DeleteBook(uint(ID)); err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent) 
}


//update function updates a book in the database
func UpdateBook(w http.ResponseWriter, r *http.Request){
	book := &models.Book{}
	utils.ParseBody(r,book)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err := strconv.ParseUint(bookId,10,64)
	if(err != nil){
		http.Error(w,"Invalid Book Id",http.StatusBadRequest)
		return
	}
	bookDetails,db := models.GetBookById(uint(ID))
	if book.Name != ""{
		bookDetails.Name = book.Name

	}
	if book.Author != ""{
		bookDetails.Author = book.Author
	}
	if book.Publication != ""{
		bookDetails.Publication = book.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	





}