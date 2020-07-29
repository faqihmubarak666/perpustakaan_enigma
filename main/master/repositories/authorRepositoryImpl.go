package repositories

import (
	"database/sql"
	"fmt"
	"gomux/main/master/constantaQuery"
	"gomux/main/master/models"
)

type AuthorRepoImpl struct {
	db *sql.DB
}

func (s AuthorRepoImpl) GetAllAuthor() ([]*models.AllAuthor, error) {
	dataAuthor := []*models.AllAuthor{}
	query := constantaQuery.GETALLAUTHOR
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		author := models.AllAuthor{}
		var err = data.Scan(&author.IdAuthor, &author.NameAuthor)
		if err != nil {
			return nil, err
		}
		dataAuthor = append(dataAuthor, &author)
	}
	fmt.Println("Successfully display Author data")
	return dataAuthor, nil
}

func InitAuthorRepoImpl(db *sql.DB) AuthorRepository {
	return &AuthorRepoImpl{db}
}
