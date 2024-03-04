package Services

import (
	"time"
	"web-api/Models/DatabaseModels"
	"web-api/Repositories/GormPsqlRepo"
)

type BlogService struct {
	blogRepo GormPsqlRepo.IBlogRepository
}

type IBlogService interface {
	CreateBlog(userId uint, title string, content string, tags []DatabaseModels.Tag) error
	GetBlogByID(id uint) (DatabaseModels.Blog, error)
	UpdateBlog(userId uint, title string, content string, tags []DatabaseModels.Tag) error
	DeleteBlog(id uint) error
}

func NewBlogService(blogRepo GormPsqlRepo.IBlogRepository) IBlogService {
	return &BlogService{blogRepo}
}

func (bs *BlogService) CreateBlog(userId uint, title string, content string, tags []DatabaseModels.Tag) error {

	blog := DatabaseModels.Blog{
		UserID:   userId,
		Title:    title,
		Content:  content,
		PostDate: time.Now().Format("2006-01-02 00:00:00"),
		Tags:     tags,
	}

	return bs.blogRepo.CreateBlog(blog)
}

func (bs *BlogService) UpdateBlog(userId uint, title string, content string, tags []DatabaseModels.Tag) error {

	blog := DatabaseModels.Blog{
		UserID:   userId,
		Title:    title,
		Content:  content,
		PostDate: time.Now().Format("2006-01-02 00:00:00"),
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
