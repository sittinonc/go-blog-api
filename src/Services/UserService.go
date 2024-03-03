package Services

import (
	"web-api/src/Models/DatabaseModels"
	"web-api/src/Repositories/GormPsqlRepo"
)

type UserService struct {
	userRepo GormPsqlRepo.UserRepository
}

type IUserService interface {
	CreateUser(username string, email string, profileImageUrl string) error
	GetUserByID(id uint) (DatabaseModels.User, error)
	UpdateUser(userId uint, username string, email string, profileImageUrl string) error
}

func NewUserService(userRepo GormPsqlRepo.UserRepository) IUserService {
	return &UserService{userRepo}
}

func (us *UserService) CreateUser(username string, email string, profileImageUrl string) error {
	user := DatabaseModels.User{
		Username:        username,
		Email:           email,
		ProfileImageUrl: profileImageUrl,
	}

	return us.userRepo.CreateUser(user)
}

func (us *UserService) GetUserByID(id uint) (DatabaseModels.User, error) {
	return us.userRepo.GetUserByID(id)
}

func (us *UserService) UpdateUser(userId uint, username string, email string, profileImageUrl string) error {
	user := DatabaseModels.User{
		Username:        username,
		Email:           email,
		ProfileImageUrl: profileImageUrl,
	}

	return us.userRepo.UpdateUser(userId, user)
}
