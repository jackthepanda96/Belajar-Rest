package data

import (
	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"github.com/jackthepanda96/Belajar-Rest.git/feature/book/data"
)

type User struct {
	ID       int         `json:"id" form:"id" gorm:"prmaryKey;autoIncrement"`
	Nama     string      `json:"nama" form:"nama" validate:"required"`
	Email    string      `json:"email" form:"email" validate:"required,email"`
	Password string      `json:"password" form:"password" validate:"required"`
	Buku     []data.Book `gorm:"foreignKey:Pemilik"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:       u.ID,
		Nama:     u.Nama,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Email = data.Email
	res.Nama = data.Nama
	res.Password = data.Password
	res.ID = data.ID
	return res
}
