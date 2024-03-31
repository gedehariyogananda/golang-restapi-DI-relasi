package domain

import (
	"test/golang/helper"
	"time"
)

type Post struct {
	Id        int64     `gorm:"primary_key"`
	User_ID   int64     `gorm:"type:int;not null"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   string    `gorm:"type:text;not null"`
	Slug      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

type PostRepository interface {
	FindAllPost() ([]Post, error)
	CreatePost(post Post) error
	FindUser(id int64) (User, error)
	FindByID(id int64) (Post, error)
}

type PostService interface {
	Store(post Post) helper.ApiResponse
	FindUser(id int64) helper.ApiResponse
}

type PostDataUser struct {
	ID        int64     `json:"id"`
	User_ID   int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	User      UserData  `json:"user"`
}
