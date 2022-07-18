package usecase

import (
	"github.com/jackthepanda96/Belajar-Rest.git/domain"
)

type bookUseCase struct {
	data domain.BookData
}

func New(model domain.BookData) domain.BookUseCase {
	return &bookUseCase{
		data: model,
	}
}
