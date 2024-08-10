package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ritankarsaha/go-bookstore/pkg/config"
)

var db *gorm.DB



//defining the book models
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}



//connecting the config package
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}



//creating new books
func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}



//getting all the books from the database
func GetAllBooks() ([]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}



//getting a single book from the databse
func GetBookById(ID uint) (*Book, error) {
	var book Book
	if err := db.Where("id = ?", ID).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}



//deleting a book from the databse
func DeleteBook(ID uint) error {
	if err := db.Where("id = ?", ID).Delete(&Book{}).Error; err != nil {
		return err
	}
	return nil
}
