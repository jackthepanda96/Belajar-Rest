package data

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penerbit string `json:"penerbit" form:"penerbit"`
	ISBN     string
	Pemilik  int
}
