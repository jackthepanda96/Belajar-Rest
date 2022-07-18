package data

import (
	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penerbit string `json:"penerbit" form:"penerbit"`
	ISBN     string
	Pemilik  int
}

func (b *Book) ToDomain() domain.Book {
	return domain.Book{
		ID:       int(b.ID),
		Judul:    b.Judul,
		Penerbit: b.Penerbit,
		ISBN:     b.ISBN,
		Pemilik:  b.Pemilik,
	}
}

func ParseToArrDomain(arr []Book) []domain.Book {
	var res []domain.Book

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}

	return res
}

func ToLocal(data domain.Book) Book {
	var res Book
	res.ID = uint(data.ID)
	res.Judul = data.Judul
	res.Penerbit = data.Penerbit
	res.Pemilik = data.Pemilik
	res.ISBN = data.ISBN
	return res
}
