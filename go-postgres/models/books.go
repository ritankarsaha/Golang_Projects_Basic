package models

import "gorm.io/gorm"

type Books struct {
	ID        uint   `json:"id" gorm:"primary_key;autoIncrement"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}