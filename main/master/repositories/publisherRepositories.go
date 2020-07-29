package repositories

import "gomux/main/master/models"

type PublisherRepository interface {
	GetAllPublisher() ([]*models.AllPublisher, error)
}
