package post

import (
	"test/golang/domain"
	"test/golang/helper"

	"github.com/gosimple/slug"
)

type service struct {
	// userRepository domain.UserRepository
	postRepository domain.PostRepository
}

func NewService(postRepository domain.PostRepository) domain.PostService {
	return &service{postRepository}
}

// Store implements domain.PostService.
func (s *service) Store(post domain.Post) helper.ApiResponse {
	post.Slug = slug.Make(post.Title)

	if err := s.postRepository.CreatePost(post); err != nil {
		return helper.ApiResponse{
			Data:    nil,
			Code:    500,
			Message: "Failed to store post",
			Status:  false,
		}
	}

	return helper.ApiResponse{
		Code:    200,
		Data:    post,
		Message: "Successfully store post",
		Status:  true,
	}
}

func (s *service) FindUser(id int64) helper.ApiResponse {

	post, err := s.postRepository.FindByID(id)

	if err != nil {
		return helper.ApiResponse{
			Data:    nil,
			Code:    404,
			Message: "Failed to find all post",
			Status:  false,
		}
	}

	user, err := s.postRepository.FindUser(post.User_ID)

	if err != nil {
		return helper.ApiResponse{
			Data:    nil,
			Code:    404,
			Message: "Failed to find user",
			Status:  false,
		}
	}

	userData := domain.UserData{
		ID:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	result := domain.PostDataUser{
		ID:        post.Id,
		User_ID:   post.User_ID,
		Title:     post.Title,
		Content:   post.Content,
		Slug:      post.Slug,
		CreatedAt: post.CreatedAt,
		User:      userData,
	}

	return helper.ApiResponse{
		Data:    result,
		Code:    200,
		Message: "Successfully find user",
		Status:  true,
	}

}
