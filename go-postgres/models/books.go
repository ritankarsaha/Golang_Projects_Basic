package models

import "gorm.io/gorm"

type Book struct {
	ID        uint   `json:"id" gorm:"primary_key;autoIncrement"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	return err
}
