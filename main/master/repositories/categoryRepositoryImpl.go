package repositories

import (
	"database/sql"
	"fmt"
	"gomux/main/master/constantaQuery"
	"gomux/main/master/models"
)

type CategoryRepoImpl struct {
	db *sql.DB
}

func (s CategoryRepoImpl) GetAllCategory() ([]*models.AllCategory, error) {
	dataCategory := []*models.AllCategory{}
	query := constantaQuery.GETALLCATEGORY
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		category := models.AllCategory{}
		var err = data.Scan(&category.IdCategory, &category.NameCategory)
		if err != nil {
			return nil, err
		}
		dataCategory = append(dataCategory, &category)
	}
	fmt.Println("Successfully display category data")
	return dataCategory, nil
}

func InitCategoryRepoImpl(db *sql.DB) CategoryRepository {
	return &CategoryRepoImpl{db}
}
