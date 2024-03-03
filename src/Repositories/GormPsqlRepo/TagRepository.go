package GormPsqlRepo

import (
	"web-api/src/Models/DatabaseModels"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

type ITagRepository interface {
	GetAllTags() error
	GetTagsByIDs(ids []uint) []DatabaseModels.Tag
	CreateTag(tag DatabaseModels.Tag) error
	UpdateTag(id uint, tag DatabaseModels.Tag) error
	DeleteTag(id uint) error
}

func NewTagRepository(db *gorm.DB) ITagRepository {
	return &TagRepository{db}
}

func (tr *TagRepository) GetAllTags() error {
	return tr.db.Find(&[]DatabaseModels.Tag{}).Error
}

func (tr *TagRepository) GetTagsByIDs(ids []uint) []DatabaseModels.Tag {
	var tags []DatabaseModels.Tag
	tr.db.Where("id IN ?", ids).Find(&tags)
	return tags

}

func (tr *TagRepository) CreateTag(tag DatabaseModels.Tag) error {
	return tr.db.Create(&tag).Error
}

func (tr *TagRepository) UpdateTag(id uint, tag DatabaseModels.Tag) error {
	var existingTag DatabaseModels.Tag
	if err := tr.db.First(&existingTag, id).Error; err != nil {
		return err
	}

	existingTag = tag
	return tr.db.Save(&existingTag).Error
}

func (tr *TagRepository) DeleteTag(id uint) error {
	return tr.db.Delete(&DatabaseModels.Tag{}, id).Error
}
