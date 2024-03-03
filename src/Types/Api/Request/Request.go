package Request

type CreateBlogRequest struct {
	UserID  uint   `json:"userID"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []uint `json:"tags"`
}

type UpdateBlogRequest struct {
	UserID  uint   `json:"userID"`
	BlogID  uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []uint `json:"tags"`
}

type CreateUserRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	ProfileImageUrl string `json:"image"`
}

type UpdateUserRequest struct {
	UserID          uint   `json:"userID"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	ProfileImageUrl string `json:"image"`
}

type CreateCommentRequest struct {
	UserID  uint   `json:"userID"`
	BlogID  uint   `json:"blogID"`
	Content string `json:"content"`
}

type UpdateCommentRequest struct {
	CommentID uint   `json:"commentID"`
	UserID    uint   `json:"userID"`
	BlogID    uint   `json:"blogID"`
	Content   string `json:"content"`
}

type CreateTagRequest struct {
	Name string `json:"name"`
}

type UpdateTagRequest struct {
	TagID uint   `json:"tagID"`
	Name  string `json:"name"`
}
