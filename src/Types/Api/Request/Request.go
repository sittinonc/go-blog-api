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
	UserID          uint   `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	ProfileImageUrl string `json:"image"`
}
