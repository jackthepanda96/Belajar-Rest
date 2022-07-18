package domain

type User struct {
	ID       int
	Nama     string
	Email    string
	Password string
}

type UserUseCase interface {
	AddUser(newUser User) (User, error)
	GetAll() ([]User, error)
}

type UserData interface {
	Insert(newUser User) User
	Update(userID int, updatedData User) User
	Delete(userID int) bool
	GetAll() []User
	GetSpecific(userID int) User
}
