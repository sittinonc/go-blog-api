package Controllers

import (
	"web-api/Models/ViewModels"
	"web-api/Services"
	"web-api/Types/Api/Request"
	"web-api/Types/Api/Response"
	"web-api/Utils"
)

type CommentController struct {
	commentService Services.CommentService
}

type ICommentController interface {
	CreateComment(request Request.CreateCommentRequest) Response.OperationResponse[error]
	GetCommentByID(id uint) Response.OperationResponse[ViewModels.Comment]
	UpdateComment(request Request.UpdateCommentRequest) Response.OperationResponse[error]
	DeleteComment(id uint) Response.OperationResponse[error]
}

func NewCommentController(commentService Services.CommentService) ICommentController {
	return &CommentController{commentService: commentService}
}

func (cc *CommentController) CreateComment(request Request.CreateCommentRequest) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}

	err := cc.commentService.CreateComment(request.UserID, request.BlogID, request.Content)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Data = err
		return response
	}

	response.Success = true
	response.Message = "Comment created successfully"
	return response
}

func (cc *CommentController) GetCommentByID(id uint) Response.OperationResponse[ViewModels.Comment] {
	response := Response.OperationResponse[ViewModels.Comment]{Success: false, Message: "", Data: ViewModels.Comment{}}
	comment, err := cc.commentService.GetCommentByID(id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "Comment retrieved successfully"
	UserViewModel := Utils.UserDbToViewModel(comment.Author)
	response.Data = ViewModels.Comment{
		CommentID: comment.ID,
		Author:    UserViewModel,
		Content:   comment.Content,
		PostDate:  comment.PostDate,
	}
	return response
}

func (cc *CommentController) UpdateComment(request Request.UpdateCommentRequest) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}

	err := cc.commentService.UpdateComment(request.CommentID, request.UserID, request.BlogID, request.Content)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "Comment updated successfully"
	return response
}

func (cc *CommentController) DeleteComment(id uint) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}

	err := cc.commentService.DeleteComment(id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "Comment deleted successfully"
	return response
}
