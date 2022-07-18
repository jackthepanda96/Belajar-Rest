package usecase

import (
	"testing"

	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	// MockTrue
	mockData := domain.User{ID: 1, Nama: "Jerry", Email: "jerry@alterra.id"}
	mockData2 := domain.User{}
	t.Run("Success case", func(t *testing.T) {
		useCase := New(&mockUserDataTrue{})
		res, err := useCase.AddUser(mockData)
		assert.Nil(t, err)                 // Apakah errornya nil
		assert.Greater(t, res.ID, 0)       // Apakah ID nya lebih besar dari 0
		assert.Equal(t, "Jerry", res.Nama) // Apakah nama yang di insertkan sama
	})

	t.Run("Cannot insert data", func(t *testing.T) {
		useCase := New(&mockUserDataFalse{})
		res, err := useCase.AddUser(mockData2)
		assert.NotNil(t, err)      // Apakah errornya tidak nil
		assert.Equal(t, 0, res.ID) // Apakah ID = 0
	})
}

type mockUserDataTrue struct{}

func (mud *mockUserDataTrue) Insert(newUser domain.User) domain.User {
	newUser.ID = 1
	return newUser
}

func (mud *mockUserDataTrue) GetAll() []domain.User {
	return []domain.User{{ID: 1, Nama: "Jerry", Email: "jerry@alterra.id", Password: "1234"}}
}

type mockUserDataFalse struct{}

func (mudf *mockUserDataFalse) Insert(newUser domain.User) domain.User {
	return domain.User{}
}

func (mudf *mockUserDataFalse) GetAll() []domain.User {
	return nil
}
