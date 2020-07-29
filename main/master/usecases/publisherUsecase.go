package usecases

import "gomux/main/master/models"

type PublisherUsecase interface {
	GetAllPublisher() ([]*models.AllPublisher, error)
}
