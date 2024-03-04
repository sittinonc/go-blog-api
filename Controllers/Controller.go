package Controllers

import "web-api/Services"

type ServerControllers struct {
	UserController    IUserController
	BlogController    IBlogController
	CommentController ICommentController
	TagController     ITagController
}

func New(userService Services.IUserService, blogService Services.IBlogService, commentService Services.ICommentService, tagService Services.ITagService) *ServerControllers {
	return &ServerControllers{
		UserController:    NewUserController(userService),
		BlogController:    NewBlogController(blogService, tagService),
		CommentController: NewCommentController(commentService),
		TagController:     NewTagController(tagService),
	}
}
