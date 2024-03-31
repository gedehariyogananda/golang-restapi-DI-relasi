package domain

import (
	"test/golang/helper"
	"time"
)

type User struct {
	Id        int64     `gorm:"primary_key"`
	Username  string    `json:"username"`
	Email     string    `gorm:"type:varchar(255);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id int64) (User, error)
	Create(user User) error
	Update(user User) error
	FindAllPost(id int64) ([]Post, error)
}

type UserService interface {
	FindAll() helper.ApiResponse
	FindByID(id int64) helper.ApiResponse
	Store(user User) helper.ApiResponse
	Update(user User) helper.ApiResponse
	FindAllPost(id int64) helper.ApiResponse
}

type UserPost struct {
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Posts    []PostDataUser `json:"posts"`
}

type UserData struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
