package usecases

import (
	"gomux/main/master/models"
	"gomux/main/master/repositories"
)

type PublisherUsecaseImpl struct {
	PublisherRepo repositories.PublisherRepository
}

func (s PublisherUsecaseImpl) GetAllPublisher() ([]*models.AllPublisher, error) {
	publisher, err := s.PublisherRepo.GetAllPublisher()
	if err != nil {
		return nil, err
	}

	return publisher, nil
}

func InitPublisherUseCase(PublisherRepo repositories.PublisherRepository) PublisherUsecase {
	return &PublisherUsecaseImpl{PublisherRepo}
}
