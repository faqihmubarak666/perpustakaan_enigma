package usecases

import "gomux/main/master/models"

type AuthorUsecase interface {
	GetAllAuthor() ([]*models.AllAuthor, error)
}
