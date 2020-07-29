package repositories

import (
	"database/sql"
	"fmt"
	"gomux/main/master/constantaQuery"
	"gomux/main/master/models"
)

type PublisherRepoImpl struct {
	db *sql.DB
}

func (s PublisherRepoImpl) GetAllPublisher() ([]*models.AllPublisher, error) {
	dataPublisher := []*models.AllPublisher{}
	query := constantaQuery.GETALLPUBLISHER
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		publisher := models.AllPublisher{}
		var err = data.Scan(&publisher.IdPublisher, &publisher.NamePublisher)
		if err != nil {
			return nil, err
		}
		dataPublisher = append(dataPublisher, &publisher)
	}
	fmt.Println("Successfully display Publisher data")
	return dataPublisher, nil
}

func InitPublisherRepoImpl(db *sql.DB) PublisherRepository {
	return &PublisherRepoImpl{db}
}
