package usecase

import (
	"github.com/hentan/final_project/internal/repository"
)

type UseCase struct {
	bookService repository.DatabaseRepo
}

func NewUseCase(bookService repository.DatabaseRepo) *UseCase {
	return &UseCase{bookService: bookService}
}

func (uc *UseCase) BookRepo() repository.DatabaseRepo {
	return uc.bookService
}
