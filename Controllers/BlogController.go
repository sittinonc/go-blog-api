package Controllers

import (
	"web-api/Models/ViewModels"
	"web-api/Services"
	"web-api/Types/Api/Request"
	"web-api/Types/Api/Response"
)

type BlogController struct {
	blogService Services.BlogService
	tagService  Services.TagService
}

type IBlogController interface {
	CreateBlog(request Request.CreateBlogRequest) Response.OperationResponse[error]
	UpdateBlog(request Request.UpdateBlogRequest) Response.OperationResponse[error]
	GetBlogByID(id uint) Response.OperationResponse[ViewModels.Blog]
	DeleteBlog(id uint) Response.OperationResponse[error]
}

func NewBlogController(blogService Services.BlogService, tagService Services.TagService) IBlogController {
	return &BlogController{blogService: blogService, tagService: tagService}
}

func (bc *BlogController) CreateBlog(request Request.CreateBlogRequest) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}
	tags, tagErr := bc.tagService.GetTagsByIDs(request.Tags)
	if tagErr != nil {
		response.Success = false
		response.Message = tagErr.Error()
		response.Data = tagErr
		return response
	}

	blogErr := bc.blogService.CreateBlog(request.UserID, request.Title, request.Content, tags)
	if blogErr != nil {
		response.Success = false
		response.Message = blogErr.Error()
		response.Data = blogErr
		return response

	}
	response.Success = true
	response.Message = "Blog created successfully"
	return response
}

func (bc *BlogController) UpdateBlog(request Request.UpdateBlogRequest) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}
	tags, tagErr := bc.tagService.GetTagsByIDs(request.Tags)
	if tagErr != nil {
		response.Success = false
		response.Message = tagErr.Error()
		response.Data = tagErr
		return response
	}

	blogRrr := bc.blogService.UpdateBlog(request.UserID, request.Title, request.Content, tags)
	if blogRrr != nil {
		response.Success = false
		response.Message = blogRrr.Error()
		return response
	}
	response.Success = true
	response.Message = "Blog updated successfully"
	return response
}

func (bc *BlogController) GetBlogByID(id uint) Response.OperationResponse[ViewModels.Blog] {
	response := Response.OperationResponse[ViewModels.Blog]{Success: false, Message: "", Data: ViewModels.Blog{}}
	blog, err := bc.blogService.GetBlogByID(id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "Blog retrieved successfully"

	UserViewModel := ViewModels.User{
		UserID:          blog.Author.ID,
		Username:        blog.Author.Username,
		Email:           blog.Author.Email,
		ProfileImageUrl: blog.Author.ProfileImageUrl,
	}
	blogViewModel := ViewModels.Blog{
		BlogID:   blog.ID,
		Title:    blog.Title,
		Content:  blog.Content,
		Author:   UserViewModel,
		PostDate: blog.PostDate,
	}
	response.Data = blogViewModel

	return response
}

func (bc *BlogController) DeleteBlog(id uint) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}
	err := bc.blogService.DeleteBlog(id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "Blog deleted successfully"
	return response
}
