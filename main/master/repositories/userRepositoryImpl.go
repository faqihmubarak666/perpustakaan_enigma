package repositories

import (
	"database/sql"
	"fmt"
	"gomux/main/master/constantaQuery"
	"gomux/main/master/models"
	"log"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (u *UserRepoImpl) Get(userIn *models.User) (bool, error) {
	query := constantaQuery.GETUSER
	row := u.db.QueryRow(query, userIn.Username, userIn.Password)
	var user = models.User{}
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		log.Print(err)
		return false, err
	}
	if userIn.Username == user.Username && userIn.Password == user.Password {
		return true, nil
	} else {
		log.Print(err)
		return false, err
	}
}

func (s UserRepoImpl) GetAllUser() ([]*models.AllUser, error) {
	dataUser := []*models.AllUser{}
	query := constantaQuery.GETALLUSER
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		user := models.AllUser{}
		var err = data.Scan(&user.IdUser, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		dataUser = append(dataUser, &user)
	}
	fmt.Println("Successfully display User data")
	return dataUser, nil
}

func InitUserRepoImpl(db *sql.DB) UserRepository {
	return &UserRepoImpl{db}
}
