package model

import (
	"log"

	"gorm.io/gorm"
)

type Login struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginModel struct {
	DB *gorm.DB
}

func (lm *LoginModel) DoLogin(data Login) User {
	var tmp User
	err := lm.DB.Where("email = ? AND password = ?", data.Email, data.Password).First(&tmp)

	if err.Error != nil {
		log.Println("Data not found")
		return User{}
	}

	return tmp
}
