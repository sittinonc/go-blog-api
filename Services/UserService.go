package Services

import (
	"web-api/Models/DatabaseModels"
	"web-api/Repositories/GormPsqlRepo"
)

type UserService struct {
	userRepo GormPsqlRepo.IUserRepository
}

type IUserService interface {
	CreateUser(username string, email string, profileImageUrl string) error
	GetUserByID(id uint) (DatabaseModels.User, error)
	UpdateUser(userId uint, username string, email string, profileImageUrl string) error
}

func NewUserService(userRepo GormPsqlRepo.IUserRepository) IUserService {
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
