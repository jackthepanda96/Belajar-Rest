package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jackthepanda96/Belajar-Rest.git/domain"
)

type bookUseCase struct {
	data domain.BookData
}

func New(model domain.BookData) domain.BookUseCase {
	return &bookUseCase{
		data: model,
	}
}

func (bu *bookUseCase) AddBook(IDUser int, newBook domain.Book) (domain.Book, error) {
	if IDUser == -1 {
		return domain.Book{}, errors.New("invalid user")
	}

	uidGen := uuid.New
	newBook.ISBN = uidGen().String()
	newBook.Pemilik = IDUser

	res := bu.data.Insert(newBook)
	if res.ID == 0 {
		return domain.Book{}, errors.New("error insert book")
	}

	return res, nil
}
func (bu *bookUseCase) GetAllBook() ([]domain.Book, error) {
	res := bu.data.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}
