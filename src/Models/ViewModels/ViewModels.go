package ViewModels

type User struct {
	UserID          uint   `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	ProfileImageUrl string `json:"image"`
}

type Blog struct {
	BlogID   uint      `json:"id"`
	Title    string    `json:"title"`
	Author   User      `json:"author"`
	Content  string    `json:"content"`
	PostDate string    `json:"postDate"`
	Tags     []Tag     `json:"tags"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	CommentID uint   `json:"id"`
	Author    User   `json:"author"`
	Content   string `json:"content"`
	PostDate  string `json:"date"`
}

type Tag struct {
	TagID uint   `json:"id"`
	Name  string `json:"name"`
}
