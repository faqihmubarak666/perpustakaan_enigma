package repositories

import "gomux/main/master/models"

type CategoryRepository interface {
	GetAllCategory() ([]*models.AllCategory, error)
}
