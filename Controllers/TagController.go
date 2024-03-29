package Controllers

import (
	"web-api/Models/ViewModels"
	"web-api/Services"
	"web-api/Types/Api/Request"
	"web-api/Types/Api/Response"
)

type TagController struct {
	tagService Services.ITagService
}

type ITagController interface {
	GetAllTags() Response.OperationResponse[[]ViewModels.Tag]
	GetTagByIDs(request Request.GetTagByIDsRequest) Response.OperationResponse[[]ViewModels.Tag]
	CreateTag(request Request.CreateTagRequest) Response.OperationResponse[error]
	UpdateTag(request Request.UpdateTagRequest) Response.OperationResponse[error]
	DeleteTag(id uint) Response.OperationResponse[error]
}

func NewTagController(tagService Services.ITagService) ITagController {
	return &TagController{tagService: tagService}
}

func (tc *TagController) GetAllTags() Response.OperationResponse[[]ViewModels.Tag] {
	response := Response.OperationResponse[[]ViewModels.Tag]{Success: false, Message: "", Data: []ViewModels.Tag{}}
	tags, err := tc.tagService.GetAllTags()
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "Tags retrieved successfully"
	TagViewModels := new([]ViewModels.Tag)

	for _, tag := range tags {
		*TagViewModels = append(*TagViewModels, ViewModels.Tag{
			TagID: tag.ID,
			Name:  tag.Name,
		})
	}
	response.Data = *TagViewModels
	return response
}

func (tc *TagController) GetTagByIDs(request Request.GetTagByIDsRequest) Response.OperationResponse[[]ViewModels.Tag] {
	response := Response.OperationResponse[[]ViewModels.Tag]{Success: false, Message: "", Data: []ViewModels.Tag{}}
	tags, err := tc.tagService.GetTagsByIDs(request.TagIDs)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "Tags retrieved successfully"
	TagViewModels := new([]ViewModels.Tag)

	for _, tag := range tags {
		*TagViewModels = append(*TagViewModels, ViewModels.Tag{
			TagID: tag.ID,
			Name:  tag.Name,
		})
	}
	response.Data = *TagViewModels
	return response
}

func (tc *TagController) CreateTag(request Request.CreateTagRequest) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}

	err := tc.tagService.CreateTag(request.Name)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Data = err
		return response
	}

	response.Success = true
	response.Message = "Tag created successfully"
	return response
}

func (tc *TagController) UpdateTag(request Request.UpdateTagRequest) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}

	err := tc.tagService.UpdateTag(request.TagID, request.Name)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "Tag updated successfully"
	return response
}

func (tc *TagController) DeleteTag(id uint) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}

	err := tc.tagService.DeleteTag(id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "Tag deleted successfully"
	return response
}
