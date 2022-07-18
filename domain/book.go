package domain

import "github.com/labstack/echo/v4"

type Book struct {
	ID       int
	Judul    string `json:"judul" form:"judul"`
	Penerbit string `json:"penerbit" form:"penerbit"`
	ISBN     string
	Pemilik  int
}

type BookHandler interface {
	InsertUser() echo.HandlerFunc
}

type BookUseCase interface {
	AddBook(IDUser int, newBook Book) (Book, error)
	GetAllBook() ([]Book, error)
	// GetMyBook(idUser int, status int) ([]Book, error)
	// UpdateBook(updateData Book) (Book, error)
	// DeleteBook(idBook int) error
}

type BookData interface {
	Insert(newBook Book) Book
	GetAll() []Book
	// Update(updatedBook Book) Book
	// Delete(idBook int) bool
	// GetSpecificBook(idUser int, ownStatus int) []Book
}
