package usecases

import (
	"gomux/main/master/models"
	"gomux/main/master/repositories"
)

type AuthorUsecaseImpl struct {
	AuthorRepo repositories.AuthorRepository
}

func (s AuthorUsecaseImpl) GetAllAuthor() ([]*models.AllAuthor, error) {
	author, err := s.AuthorRepo.GetAllAuthor()
	if err != nil {
		return nil, err
	}

	return author, nil
}

func InitAuthorUseCase(AuthorRepo repositories.AuthorRepository) AuthorUsecase {
	return &AuthorUsecaseImpl{AuthorRepo}
}
