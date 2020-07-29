package usecases

import (
	"gomux/main/master/models"
	"gomux/main/master/repositories"
)

type CategoryUsecaseImpl struct {
	CategoryRepo repositories.CategoryRepository
}

func (s CategoryUsecaseImpl) GetAllCategory() ([]*models.AllCategory, error) {
	category, err := s.CategoryRepo.GetAllCategory()
	if err != nil {
		return nil, err
	}

	return category, nil
}

func InitCategoryUseCase(CategoryRepo repositories.CategoryRepository) CategoryUsecase {
	return &CategoryUsecaseImpl{CategoryRepo}
}
