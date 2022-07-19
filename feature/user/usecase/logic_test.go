package usecase

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"github.com/jackthepanda96/Belajar-Rest.git/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddUser(t *testing.T) {
	repo := new(mocks.UserData)

	mockData := domain.User{Nama: "Jerry", Email: "jerry@alterra.id", Password: "123"}
	// mockData2 := domain.User{}
	returnData := mockData
	returnData.ID = 1
	returnData.Password = "$2a$10$OqHN2OI/X2g8c5on5JV33.m0vLv4U5nhniXpb.hu2ddcSSj/nZMFq"
	t.Run("Success case", func(t *testing.T) {
		// useCase := New(&mockUserDataTrue{})
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.AddUser(mockData)
		assert.Nil(t, err)                 // Apakah errornya nil
		assert.Greater(t, res.ID, 0)       // Apakah ID nya lebih besar dari 0
		assert.Equal(t, "Jerry", res.Nama) // Apakah nama yang di insertkan sama
		assert.Equal(t, "jerry@alterra.id", res.Email)
		assert.Equal(t, "$2a$10$OqHN2OI/X2g8c5on5JV33.m0vLv4U5nhniXpb.hu2ddcSSj/nZMFq", res.Password, "Password tidak sesuai")
		repo.AssertExpectations(t)
	})

	t.Run("Validator error", func(t *testing.T) {
		// useCase := New(&mockUserDataTrue{})
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.AddUser(domain.User{})
		assert.EqualError(t, err, "error") // Apakah errornya nil
		assert.Greater(t, res.ID, 0)       // Apakah ID nya lebih besar dari 0
		assert.Equal(t, "", res.Nama)      // Apakah nama yang di insertkan sama
		assert.Equal(t, "", res.Email)
		assert.Equal(t, "", res.Password, "Password tidak sesuai")
		repo.AssertExpectations(t)
	})

	// t.Run("Duplicated Data", func(t *testing.T) {
	// 	repo.On("Insert", mock.Anything).Return(domain.User{}, gorm.ErrRegistered).Once()
	// 	useCase := New(repo)
	// 	res, err := useCase.AddUser(returnData)
	// 	assert.NotNil(t, err) // Apakah errornya tidak nil
	// 	assert.EqualError(t, err, gorm.ErrRegistered.Error())
	// 	assert.Equal(t, 0, res.ID)        // Apakah ID = 0
	// 	assert.Equal(t, "", res.Nama)     // Apakah ID = 0
	// 	assert.Equal(t, "", res.Email)    // Apakah ID = 0
	// 	assert.Equal(t, "", res.Password) // Apakah ID = 0
	// 	repo.AssertExpectations(t)
	// })

	// t.Run("Error from server", func(t *testing.T) {
	// 	repo.On("Insert", mock.Anything).Return(domain.User{}, gorm.ErrInvalidValueOfLength).Once()
	// 	useCase := New(repo)
	// 	returnData.Email = "123aoeijakdngnsvbsnzoczbjfakdjfoadijfoangnbcoloijapdfaposdjfpk"
	// 	res, err := useCase.AddUser(returnData)
	// 	assert.NotNil(t, err) // Apakah errornya tidak nil
	// 	assert.EqualError(t, err, gorm.ErrInvalidValueOfLength.Error())
	// 	assert.Equal(t, 0, res.ID)        // Apakah ID = 0
	// 	assert.Equal(t, "", res.Nama)     // Apakah ID = 0
	// 	assert.Equal(t, "", res.Email)    // Apakah ID = 0
	// 	assert.Equal(t, "", res.Password) // Apakah ID = 0
	// 	repo.AssertExpectations(t)
	// })
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.UserData) // Menggunakan mock object yang sudah dibuat

	t.Run("Success get all", func(t *testing.T) {
		repo.On("GetAll").Return([]domain.User{{ID: 1, Nama: "Jerry", Email: "jerry@alterra.id", Password: "1234"}}, nil).Once()
		// usecase := New(&mockUserDataTrue{})
		usecase := New(repo, validator.New())
		res, err := usecase.GetAll()
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Greater(t, res[0].ID, 0)
		repo.AssertExpectations(t)
	})

	t.Run("Error not found", func(t *testing.T) {
		repo.On("GetAll").Return(nil, gorm.ErrRecordNotFound).Once()
		// usecase := New(&mockUserDataFalse{})
		usecase := New(repo, validator.New())
		res, err := usecase.GetAll()
		assert.NotNil(t, err)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error cannot retrieve data", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("cannot retrieve data")).Once()
		// usecase := New(&mockUserDataFalse{})
		usecase := New(repo, validator.New())
		res, err := usecase.GetAll()
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.New("error when retrieve data").Error())
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}
