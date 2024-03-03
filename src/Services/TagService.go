package Services

import (
	"web-api/src/Models/DatabaseModels"

	"web-api/src/Repositories/GormPsqlRepo"
)

type TagService struct {
	tagRepo GormPsqlRepo.TagRepository
}

type ITagService interface {
	GetAllTags() ([]DatabaseModels.Tag, error)
	GetTagsByIDs(ids []uint) ([]DatabaseModels.Tag, error)
	CreateTag(name string) error
	UpdateTag(id uint, name string) error
	DeleteTag(id uint) error
}

func NewTagService(tagRepo GormPsqlRepo.TagRepository) ITagService {
	return &TagService{tagRepo}
}

func (ts *TagService) GetAllTags() ([]DatabaseModels.Tag, error) {
	return ts.tagRepo.GetAllTags()
}

func (ts *TagService) GetTagsByIDs(ids []uint) ([]DatabaseModels.Tag, error) {
	return ts.tagRepo.GetTagsByIDs(ids)
}

func (ts *TagService) CreateTag(name string) error {
	tag := DatabaseModels.Tag{
		Name: name,
	}

	return ts.tagRepo.CreateTag(tag)
}

func (ts *TagService) UpdateTag(id uint, name string) error {
	tag := DatabaseModels.Tag{
		Name: name,
	}

	return ts.tagRepo.UpdateTag(id, tag)
}

func (ts *TagService) DeleteTag(id uint) error {
	return ts.tagRepo.DeleteTag(id)
}
