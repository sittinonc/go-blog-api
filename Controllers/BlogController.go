package Controllers

import (
	"web-api/Models/ViewModels"
	"web-api/Services"
	"web-api/Types/Api/Request"
	"web-api/Types/Api/Response"
)

type BlogController struct {
	blogService Services.BlogService
}

type IBlogController interface {
	CreateBlog(request Request.CreateBlogRequest) Response.OperationResponse[error]
	UpdateBlog(request Request.UpdateBlogRequest) Response.OperationResponse[error]
	GetBlogByID(id uint) Response.OperationResponse[ViewModels.Blog]
	DeleteBlog(id uint) Response.OperationResponse[error]
}

func NewBlogController(blogService Services.BlogService) IBlogController {
	return &BlogController{blogService}
}

func (bc *BlogController) CreateBlog(request Request.CreateBlogRequest) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}
	err := bc.blogService.CreateBlog(request.UserID, request.Title, request.Content, request.Tags)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Data = err
		return response

	}
	response.Success = true
	response.Message = "Blog created successfully"
	return response
}

func (bc *BlogController) UpdateBlog(request Request.UpdateBlogRequest) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}
	err := bc.blogService.UpdateBlog(request.UserID, request.Title, request.Content, request.Tags)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
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