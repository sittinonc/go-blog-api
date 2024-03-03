package Services

import (
	"web-api/src/Models/DatabaseModels"
	"web-api/src/Repositories/GormPsqlRepo"
)

type CommentService struct {
	commentRepo GormPsqlRepo.CommentRepository
}

type ICommentService interface {
	CreateComment(comment DatabaseModels.Comment) error
	GetCommentByID(id uint) (DatabaseModels.Comment, error)
	UpdateComment(comment DatabaseModels.Comment) error
	DeleteComment(id uint) error
}

func NewCommentService(commentRepo GormPsqlRepo.CommentRepository) ICommentService {
	return &CommentService{commentRepo}
}

func (cs *CommentService) CreateComment(comment DatabaseModels.Comment) error {
	return cs.commentRepo.CreateComment(comment)
}

func (cs *CommentService) GetCommentByID(id uint) (DatabaseModels.Comment, error) {
	return cs.commentRepo.GetCommentByID(id)
}

func (cs *CommentService) UpdateComment(comment DatabaseModels.Comment) error {
	return cs.commentRepo.UpdateComment(comment)
}

func (cs *CommentService) DeleteComment(id uint) error {
	return cs.commentRepo.DeleteComment(id)
}
