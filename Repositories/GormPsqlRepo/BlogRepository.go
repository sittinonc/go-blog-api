package GormPsqlRepo

import (
	"web-api/Models/DatabaseModels"

	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

type IBlogRepository interface {
	CreateBlog(blog DatabaseModels.Blog) error
	GetBlogByID(id uint) (DatabaseModels.Blog, error)
	UpdateBlog(blog DatabaseModels.Blog) error
	DeleteBlog(id uint) error
}

func NewBlogRepository(db *gorm.DB) IBlogRepository {
	return &BlogRepository{db}
}

func (br *BlogRepository) CreateBlog(blog DatabaseModels.Blog) error {
	return br.db.Create(&blog).Error
}

func (br *BlogRepository) GetBlogByID(id uint) (DatabaseModels.Blog, error) {
	var blog DatabaseModels.Blog
	err := br.db.Preload("Comments").Preload("Tags").Preload("User").First(&blog, id).Error
	return blog, err
}

func (br *BlogRepository) UpdateBlog(blog DatabaseModels.Blog) error {
	return br.db.Save(&blog).Error
}

func (br *BlogRepository) DeleteBlog(id uint) error {
	return br.db.Delete(&DatabaseModels.Blog{}, id).Error
}
