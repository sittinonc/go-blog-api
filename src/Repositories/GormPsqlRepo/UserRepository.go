package GormPsqlRepo

import (
	"web-api/src/Models/DatabaseModels"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	CreateUser(user DatabaseModels.User) error
	GetUserByID(id uint) (DatabaseModels.User, error)
	UpdateUser(id uint, user DatabaseModels.User) error
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user DatabaseModels.User) error {
	return ur.db.Create(&user).Error
}

func (ur *UserRepository) GetUserByID(id uint) (DatabaseModels.User, error) {
	var user DatabaseModels.User
	err := ur.db.First(&user, id).Error
	return user, err
}

func (ur *UserRepository) UpdateUser(id uint, user DatabaseModels.User) error {
	var existingUser DatabaseModels.User
	err := ur.db.First(&existingUser, id).Error
	if err != nil {
		return err
	}

	return ur.db.Save(&existingUser).Error
}
