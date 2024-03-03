package Services

import (
	"web-api/Models/DatabaseModels"
	"web-api/Repositories/GormPsqlRepo"
)

type BlogService struct {
	blogRepo GormPsqlRepo.BlogRepository
}

type IBlogService interface {
	CreateBlog(userId uint, title string, content string, tags []uint) error
	GetBlogByID(id uint) (DatabaseModels.Blog, error)
	UpdateBlog(userId uint, title string, content string, tags []uint) error
	DeleteBlog(id uint) error
}

func NewBlogService(blogRepo GormPsqlRepo.BlogRepository) IBlogService {
	return &BlogService{blogRepo}
}

func (bs *BlogService) CreateBlog(userId uint, title string, content string, tags []uint) error {
	// mock data
	mockTags := []DatabaseModels.Tag{}
	mockPostDate := "2021-01-01"

	blog := DatabaseModels.Blog{
		UserID:   userId,
		Title:    title,
		Content:  content,
		PostDate: mockPostDate,
		Tags:     mockTags,
	}

	return bs.blogRepo.CreateBlog(blog)
}

func (bs *BlogService) GetBlogByID(id uint) (DatabaseModels.Blog, error) {
	return bs.blogRepo.GetBlogByID(id)
}

func (bs *BlogService) UpdateBlog(userId uint, title string, content string, tags []uint) error {
	// mock data
	mockTags := []DatabaseModels.Tag{}
	mockPostDate := "2021-01-01"

	blog := DatabaseModels.Blog{
		UserID:   userId,
		Title:    title,
		Content:  content,
		PostDate: mockPostDate,
		Tags:     mockTags,
	}

	return bs.blogRepo.UpdateBlog(blog)
}

func (bs *BlogService) DeleteBlog(id uint) error {
	return bs.blogRepo.DeleteBlog(id)
}
