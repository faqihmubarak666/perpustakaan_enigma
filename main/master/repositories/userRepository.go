package repositories

import "gomux/main/master/models"

type UserRepository interface {
	Get(*models.User) (bool, error)
	GetAllUser() ([]*models.AllUser, error)
}
