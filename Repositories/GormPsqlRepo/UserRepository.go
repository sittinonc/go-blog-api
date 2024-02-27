package GormPsqlRepo

import (
	"web-api/Models/DatabaseModels"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	CreateUser(user DatabaseModels.User) error
	GetUserByID(id uint) (DatabaseModels.User, error)
	UpdateUser(user DatabaseModels.User) error
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

func (ur *UserRepository) UpdateUser(user DatabaseModels.User) error {
	return ur.db.Save(&user).Error
}
