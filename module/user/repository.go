package user

import (
	"test/golang/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.UserRepository {
	return &repository{db}
}

// FindAll implements domain.UserRepository.
func (r *repository) FindAll() ([]domain.User, error) {
	var users []domain.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Create implements domain.UserRepository.
func (r *repository) Create(user domain.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// FindByID implements domain.UserRepository.
func (r *repository) FindByID(id int64) (domain.User, error) {
	var user domain.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// Update implements domain.UserRepository.
func (r *repository) Update(user domain.User) error {
	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

// FindAllPost implements domain.UserRepository.
func (r *repository) FindAllPost(id int64) ([]domain.Post, error) {
	var posts []domain.Post

	if err := r.db.Where("user_id = ?", id).Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil

}
