package delivery

import "github.com/jackthepanda96/Belajar-Rest.git/domain"

type BookInsertRequest struct {
	Judul    string `json:"judul"`
	Penerbit string `json:"penerbit"`
}

func (bi *BookInsertRequest) ToDomain() domain.Book {
	return domain.Book{
		Judul:    bi.Judul,
		Penerbit: bi.Penerbit,
	}
}
