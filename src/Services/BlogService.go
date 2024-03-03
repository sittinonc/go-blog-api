package Services

import (
	"web-api/src/Models/DatabaseModels"
	"web-api/src/Repositories/GormPsqlRepo"
)

type BlogService struct {
	blogRepo GormPsqlRepo.BlogRepository
}

type IBlogService interface {
	CreateBlog(userId uint, title string, content string, tags []DatabaseModels.Tag) error
	GetBlogByID(id uint) (DatabaseModels.Blog, error)
	UpdateBlog(userId uint, title string, content string, tags []DatabaseModels.Tag) error
	DeleteBlog(id uint) error
}

func NewBlogService(blogRepo GormPsqlRepo.BlogRepository) IBlogService {
	return &BlogService{blogRepo}
}

func (bs *BlogService) CreateBlog(userId uint, title string, content string, tags []DatabaseModels.Tag) error {
	// mock data
	mockPostDate := "2021-01-01"

	blog := DatabaseModels.Blog{
		UserID:   userId,
		Title:    title,
		Content:  content,
		PostDate: mockPostDate,
		Tags:     tags,
	}

	return bs.blogRepo.CreateBlog(blog)
}

func (bs *BlogService) UpdateBlog(userId uint, title string, content string, tags []DatabaseModels.Tag) error {
	// mock data
	mockPostDate := "2021-01-01"

	blog := DatabaseModels.Blog{
		UserID:   userId,
		Title:    title,
		Content:  content,
		PostDate: mockPostDate,
		Tags:     tags,
	}

	return bs.blogRepo.UpdateBlog(blog)
}

func (bs *BlogService) GetBlogByID(id uint) (DatabaseModels.Blog, error) {
	return bs.blogRepo.GetBlogByID(id)
}

func (bs *BlogService) DeleteBlog(id uint) error {
	return bs.blogRepo.DeleteBlog(id)
}
