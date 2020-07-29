package repositories

import "gomux/main/master/models"

type AuthorRepository interface {
	GetAllAuthor() ([]*models.AllAuthor, error)
}
