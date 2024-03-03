package DatabaseModels

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string `gorm:"unique"`
	Email           string `gorm:"unique"`
	ProfileImageUrl string
	HashedPassword  string
	Salt            string
	Blogs           []Blog
}

type Blog struct {
	gorm.Model
	Title    string
	Content  string
	PostDate string
	UserID   uint
	Author   User `gorm:"foreignKey:UserID"`
	Comments []Comment
	Tags     []Tag `gorm:"many2many:blog_tags;"`
}

type Comment struct {
	gorm.Model
	Content  string
	PostDate string
	UserID   uint
	Author   User `gorm:"foreignKey:UserID"`
	BlogID   uint
	Blog     Blog `gorm:"foreignKey:BlogID"`
}

type Tag struct {
	gorm.Model
	Name string `gorm:"unique"`
}
