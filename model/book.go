package model

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penerbit string `json:"penerbit" form:"penerbit"`
	ISBN     string
	Pemilik  int `json:"pemilik" form:"pemilik"`
}

type BookModel struct {
	DB *gorm.DB
}

func (bm *BookModel) Insert(newBook Book) Book {
	str := uuid.New()
	newBook.ISBN = str.String()
	err := bm.DB.Create(&newBook).Error
	if err != nil {
		log.Println("Cannot create object", err.Error())
		return Book{}
	}

	return newBook
}
func (bm *BookModel) GetAll() []Book {
	var tmp []Book
	err := bm.DB.Find(&tmp).Error

	if err != nil {
		log.Println("Cannot retrive object", err.Error())
		return nil
	}
	return tmp
}
