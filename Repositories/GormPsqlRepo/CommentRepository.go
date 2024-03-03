package GormPsqlRepo

import (
	"web-api/Models/DatabaseModels"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

type ICommentRepository interface {
	CreateComment(comment DatabaseModels.Comment) error
	GetCommentByID(id uint) (DatabaseModels.Comment, error)
	UpdateComment(id uint, comment DatabaseModels.Comment) error
	DeleteComment(id uint) error
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &CommentRepository{db}
}

func (cr *CommentRepository) CreateComment(comment DatabaseModels.Comment) error {
	return cr.db.Create(&comment).Error
}

func (cr *CommentRepository) GetCommentByID(id uint) (DatabaseModels.Comment, error) {
	var comment DatabaseModels.Comment
	err := cr.db.First(&comment, id).Error
	return comment, err
}

func (cr *CommentRepository) UpdateComment(id uint, comment DatabaseModels.Comment) error {
	var existingComment DatabaseModels.Comment
	if err := cr.db.First(&existingComment, id).Error; err != nil {
		return err
	}

	existingComment = comment
	return cr.db.Save(&existingComment).Error
}

func (cr *CommentRepository) DeleteComment(id uint) error {
	return cr.db.Delete(&DatabaseModels.Comment{}, id).Error
}
