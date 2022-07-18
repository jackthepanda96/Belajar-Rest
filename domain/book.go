package domain

type Book struct {
	Judul    string `json:"judul" form:"judul"`
	Penerbit string `json:"penerbit" form:"penerbit"`
	ISBN     string
	Pemilik  int
}

type BookUseCase interface {
}

type BookData interface {
}
