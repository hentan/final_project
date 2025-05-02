package provider

import (
	"github.com/hentan/final_project/internal/repository"
	"github.com/hentan/final_project/internal/usecase"
)

type UseCaseProvider struct {
	BookUseCase *usecase.UseCase
}

func NewUseCaseProvider(repo repository.DatabaseRepo) *UseCaseProvider {
	return &UseCaseProvider{
		BookUseCase: usecase.NewUseCase(repo),
	}
}
