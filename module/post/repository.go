package post

import (
	"test/golang/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.PostRepository {
	return &repository{db}
}

// CreatePost implements domain.PostRepository.
func (r *repository) CreatePost(post domain.Post) error {
	if err := r.db.Create(&post).Error; err != nil {
		return err
	}

	return nil
}

// FindUser implements domain.PostRepository.
func (r *repository) FindUser(id int64) (domain.User, error) {
	var user domain.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// FindByID implements domain.PostRepository.
func (r *repository) FindByID(id int64) (domain.Post, error) {
	var post domain.Post

	if err := r.db.Where("id = ?", id).First(&post).Error; err != nil {
		return post, err
	}

	return post, nil
}

// FindAllPost implements domain.PostRepository.
func (r *repository) FindAllPost() ([]domain.Post, error) {
	var posts []domain.Post

	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
