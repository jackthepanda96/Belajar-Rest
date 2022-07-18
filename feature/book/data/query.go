package data

import (
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
