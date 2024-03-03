package Services

import (
	"web-api/Models/DatabaseModels"
	"web-api/Repositories/GormPsqlRepo"
)

type TagService struct {
	tagRepo GormPsqlRepo.TagRepository
}

type ITagService interface {
	GetAllTags() error
	CreateTag(tag DatabaseModels.Tag) error
	UpdateTag(tag DatabaseModels.Tag) error
	DeleteTag(id uint) error
}

func NewTagService(tagRepo GormPsqlRepo.TagRepository) ITagService {
	return &TagService{tagRepo}
}

func (ts *TagService) GetAllTags() error {
	return ts.tagRepo.GetAllTags()
}

func (ts *TagService) CreateTag(tag DatabaseModels.Tag) error {
	return ts.tagRepo.CreateTag(tag)
}

func (ts *TagService) UpdateTag(tag DatabaseModels.Tag) error {
	return ts.tagRepo.UpdateTag(tag)
}

func (ts *TagService) DeleteTag(id uint) error {
	return ts.tagRepo.DeleteTag(id)
}
