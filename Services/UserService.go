package services

import (
	"web-api/Models/DatabaseModels"
	"web-api/Repositories/GormPsqlRepo"
)

type UserService struct {
	userRepo GormPsqlRepo.UserRepository
}

type IUserService interface {
	CreateUser(user DatabaseModels.User) error
	GetUserByID(id uint) (DatabaseModels.User, error)
	UpdateUser(user DatabaseModels.User) error
}

func NewUserService(userRepo GormPsqlRepo.UserRepository) IUserService {
	return &UserService{userRepo}
}

func (us *UserService) CreateUser(user DatabaseModels.User) error {
	return us.userRepo.CreateUser(user)
}

func (us *UserService) GetUserByID(id uint) (DatabaseModels.User, error) {
	return us.userRepo.GetUserByID(id)
}

func (us *UserService) UpdateUser(user DatabaseModels.User) error {
	return us.userRepo.UpdateUser(user)
}
