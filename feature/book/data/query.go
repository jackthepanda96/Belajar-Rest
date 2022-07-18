package data

import (
	"log"

	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.BookData {
	return &bookData{
		db: DB,
	}
}

func (bd *bookData) Insert(newBook domain.Book) domain.Book {
	cnv := ToLocal(newBook)
	err := bd.db.Create(&cnv)
	if err.Error != nil {
		log.Println("Cannot insert data", err.Error.Error())
		return domain.Book{}
	}

	return cnv.ToDomain()
}
func (bd *bookData) GetAll() []domain.Book {
	var data []Book
	err := bd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArrDomain(data)
}

// func (bd *bookData) Update(updatedBook domain.Book) domain.Book {

// }
// func (bd *bookData) Delete(idBook int) bool                                  {}
// func (bd *bookData) GetSpecificBook(idUser int, ownStatus int) []domain.Book {}
