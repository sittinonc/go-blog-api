package services

import (
	"web-api/Models/DatabaseModels"
	"web-api/Repositories/GormPsqlRepo"
)

type BlogService struct {
	blogRepo GormPsqlRepo.BlogRepository
}

type IBlogService interface {
	CreateBlog(blog DatabaseModels.Blog) error
	GetBlogByID(id uint) (DatabaseModels.Blog, error)
	UpdateBlog(blog DatabaseModels.Blog) error
	DeleteBlog(id uint) error
}

func NewBlogService(blogRepo GormPsqlRepo.BlogRepository) IBlogService {
	return &BlogService{blogRepo}
}

func (bs *BlogService) CreateBlog(blog DatabaseModels.Blog) error {
	return bs.blogRepo.CreateBlog(blog)
}

func (bs *BlogService) GetBlogByID(id uint) (DatabaseModels.Blog, error) {
	return bs.blogRepo.GetBlogByID(id)
}

func (bs *BlogService) UpdateBlog(blog DatabaseModels.Blog) error {
	return bs.blogRepo.UpdateBlog(blog)
}

func (bs *BlogService) DeleteBlog(id uint) error {
	return bs.blogRepo.DeleteBlog(id)
}
