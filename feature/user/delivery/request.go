package delivery

import "github.com/jackthepanda96/Belajar-Rest.git/domain"

type InsertFormat struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		Nama:     i.Nama,
		Email:    i.Email,
		Password: i.Password,
	}
}

type LoginFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
