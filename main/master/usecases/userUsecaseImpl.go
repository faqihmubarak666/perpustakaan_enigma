package usecases

import (
	"gomux/main/master/models"
	"gomux/main/master/repositories"
)

type UserUsecaseImpl struct {
	userRepo repositories.UserRepository
}

func (u *UserUsecaseImpl) GetUser(data *models.User) (bool, error) {
	isValid, err := u.userRepo.Get(data)
	if err != nil {
		return false, err
	}
	return isValid, nil
}
func (s UserUsecaseImpl) GetAllUser() ([]*models.AllUser, error) {
	user, err := s.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func InitUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &UserUsecaseImpl{userRepo}
}
