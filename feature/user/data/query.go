package data

import (
	"log"

	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

func (ud *userData) Insert(newUser domain.User) domain.User {
	var cnv = FromModel(newUser)
	err := ud.db.Create(&cnv).Error
	if err != nil {
		log.Println("Cannot create object", err.Error())
		return domain.User{}
	}

	return cnv.ToModel()
}
func (ud *userData) Update(userID int, updatedData domain.User) domain.User {
	var cnv = FromModel(updatedData)
	err := ud.db.Model(&User{}).Where("ID = ?", userID).Updates(updatedData).Error
	if err != nil {
		log.Println("Cannot update data", err.Error())
		return domain.User{}
	}
	updatedData.ID = userID
	return cnv.ToModel()
}
func (ud *userData) Delete(userID int) bool {
	res := ud.db.Where("ID = ?", userID).Delete(&User{})
	if res.Error != nil {
		log.Println("Cannot delete data", res.Error.Error())
		return false
	}

	if res.RowsAffected < 1 {
		log.Println("No data deleted", res.Error.Error())
		return false
	}

	return true
}
func (ud *userData) GetAll() []domain.User {
	var tmp []User
	err := ud.db.Find(&tmp).Error

	if err != nil {
		log.Println("Cannot retrive object", err.Error())
		return nil
	}
	return ParseToArr(tmp)
}
func (ud *userData) GetSpecific(userID int) domain.User {
	var tmp User
	err := ud.db.Where("ID = ?", userID).First(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return domain.User{}
	}

	return tmp.ToModel()
}
