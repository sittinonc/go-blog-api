package Services

import (
	"time"
	"web-api/src/Models/DatabaseModels"
	"web-api/src/Repositories/GormPsqlRepo"
)

type CommentService struct {
	commentRepo GormPsqlRepo.CommentRepository
}

type ICommentService interface {
	CreateComment(userId uint, blogId uint, content string) error
	GetCommentByID(id uint) (DatabaseModels.Comment, error)
	UpdateComment(id uint, userId uint, blogId uint, content string) error
	DeleteComment(id uint) error
}

func NewCommentService(commentRepo GormPsqlRepo.CommentRepository) ICommentService {
	return &CommentService{commentRepo}
}

func (cs *CommentService) CreateComment(userId uint, blogId uint, content string) error {
	comment := DatabaseModels.Comment{
		Content:  content,
		UserID:   userId,
		BlogID:   blogId,
		PostDate: time.Now().Format("2006-01-02 00:00:00"),
	}
	return cs.commentRepo.CreateComment(comment)
}

func (cs *CommentService) GetCommentByID(id uint) (DatabaseModels.Comment, error) {
	return cs.commentRepo.GetCommentByID(id)
}

func (cs *CommentService) UpdateComment(id uint, userId uint, blogId uint, content string) error {
	comment := DatabaseModels.Comment{
		Content:  content,
		UserID:   userId,
		BlogID:   blogId,
		PostDate: time.Now().Format("2006-01-02 00:00:00"),
	}
	return cs.commentRepo.UpdateComment(id, comment)
}

func (cs *CommentService) DeleteComment(id uint) error {
	return cs.commentRepo.DeleteComment(id)
}
