package usecase

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"github.com/jackthepanda96/Belajar-Rest.git/feature/user/data"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUseCase struct {
	userData domain.UserData
	validate *validator.Validate
}

func New(ud domain.UserData, v *validator.Validate) domain.UserUseCase {
	return &userUseCase{
		userData: ud,
		validate: v,
	}
}

func (ud *userUseCase) AddUser(newUser domain.User) (domain.User, error) {
	var cnv = data.FromModel(newUser)
	err := ud.validate.Struct(cnv)
	if err != nil {
		log.Println("Validation errror : ", err.Error())
		return domain.User{}, err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrpt password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	inserted, err := ud.userData.Insert(newUser)

	if err != nil {
		log.Println("User Usecase", err.Error())
		return domain.User{}, err
	}

	if inserted.ID == 0 {
		return domain.User{}, errors.New("cannot insert data")
	}

	return inserted, nil
}
func (ud *userUseCase) GetAll() ([]domain.User, error) {
	data, err := ud.userData.GetAll()

	if err == gorm.ErrRecordNotFound {
		log.Println("User Usecase", err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Println("User Usecase", err.Error())
		return nil, errors.New("error when retrieve data")
	}

	return data, nil
}
func (ud *userUseCase) GetProfile(id int) (domain.User, error) {
	data, err := ud.userData.GetSpecific(id)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data not found")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}

	return data, nil
}

// func (ud *userUseCase) GetSpecific() {

// }

// func (ud *userUseCase) AddUser() {

// }
// func (ud *userUseCase) AddUser() {

// }
