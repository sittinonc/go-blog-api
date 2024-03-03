package Controllers

import (
	"web-api/src/Models/ViewModels"
	"web-api/src/Services"
	"web-api/src/Types/Api/Request"
	"web-api/src/Types/Api/Response"
)

type UserController struct {
	userService Services.UserService
}

type IUserController interface {
	CreateUser(request Request.CreateUserRequest) Response.OperationResponse[error]
	UpdateUser(request Request.UpdateUserRequest) Response.OperationResponse[error]
	GetUserByID(id uint) Response.OperationResponse[ViewModels.User]
}

func NewUserController(userService Services.UserService) IUserController {
	return &UserController{userService: userService}
}

func (uc *UserController) CreateUser(request Request.CreateUserRequest) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}

	err := uc.userService.CreateUser(request.Username, request.Email, request.ProfileImageUrl)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Data = err
		return response
	}

	response.Success = true
	response.Message = "User created successfully"
	return response
}

func (uc *UserController) UpdateUser(request Request.UpdateUserRequest) Response.OperationResponse[error] {
	response := Response.OperationResponse[error]{Success: false, Message: "", Data: nil}

	err := uc.userService.UpdateUser(request.UserID, request.Username, request.Email, request.ProfileImageUrl)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "User updated successfully"
	return response
}

func (uc *UserController) GetUserByID(id uint) Response.OperationResponse[ViewModels.User] {
	response := Response.OperationResponse[ViewModels.User]{Success: false, Message: "", Data: ViewModels.User{}}
	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return response
	}
	response.Success = true
	response.Message = "User retrieved successfully"
	response.Data = ViewModels.User{
		UserID:          user.ID,
		Username:        user.Username,
		Email:           user.Email,
		ProfileImageUrl: user.ProfileImageUrl,
	}
	return response
}
