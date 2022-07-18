package usecase

import (
	"errors"
	"log"

	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData domain.UserData
}

func New(ud domain.UserData) domain.UserUseCase {
	return &userUseCase{
		userData: ud,
	}
}

func (ud *userUseCase) AddUser(newUser domain.User) (domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrpt password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	inserted := ud.userData.Insert(newUser)
	if inserted.ID == 0 {
		return domain.User{}, errors.New("cannot insert data")
	}

	return inserted, nil
}
func (ud *userUseCase) GetAll() ([]domain.User, error) {
	data := ud.userData.GetAll()

	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	return data, nil
}

// func (ud *userUseCase) GetSpecific() {

// }

// func (ud *userUseCase) AddUser() {

// }
// func (ud *userUseCase) AddUser() {

// }
