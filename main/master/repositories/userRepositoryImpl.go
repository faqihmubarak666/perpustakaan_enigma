package repositories

import (
	"database/sql"
	"gomux/main/master/constantaQuery"
	"gomux/main/master/models"
	"log"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (u *UserRepoImpl) Get(userIn *models.User) (bool, error) {
	var user = models.User{}
	query := constantaQuery.GETUSER
	row := u.db.QueryRow(query, userIn.Username, userIn.Password)
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

func InitUserRepoImpl(db *sql.DB) UserRepository {
	return &UserRepoImpl{db}
}
