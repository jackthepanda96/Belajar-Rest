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
	GetProfile(id int) (User, error)
}

type UserData interface {
	Insert(newUser User) (User, error)
	GetAll() ([]User, error)
	GetSpecific(userID int) (User, error)
	// Update(userID int, updatedData User) User
	// Delete(userID int) bool
}
