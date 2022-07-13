package model

type User struct {
	ID       int    `json:"id" form:"id" gorm:"prmaryKey;autoIncrement"`
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
