package usecases

import "gomux/main/master/models"

type CategoryUsecase interface {
	GetAllCategory() ([]*models.AllCategory, error)
}
