package usecases

import "gomux/main/master/models"

type UserUsecase interface {
	GetUser(*models.User) (bool, error)
	GetAllUser() ([]*models.AllUser, error)
}
