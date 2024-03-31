package user

import (
	"net/http"
	"test/golang/domain"
	"test/golang/helper"
)

type service struct {
	userRepository domain.UserRepository
	// postRepository domain.PostRepository
}

func NewService(userRepository domain.UserRepository) domain.UserService {
	return &service{userRepository}
}

// FindAll implements domain.UserService.
func (s *service) FindAll() helper.ApiResponse {
	dataset, err := s.userRepository.FindAll()

	if err != nil {
		return helper.ApiResponse{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed to find all",
			Status:  false,
		}
	}

	return helper.ApiResponse{
		Data:    dataset,
		Code:    http.StatusOK,
		Message: "Successfully found all",
		Status:  true,
	}
}

// FindByID implements domain.UserService.
func (s *service) FindByID(id int64) helper.ApiResponse {
	dataset, err := s.userRepository.FindByID(id)

	if err != nil {
		return helper.ApiResponse{
			Data:    nil,
			Code:    http.StatusNotFound,
			Message: "Failed to find by ID",
			Status:  false,
		}
	}

	return helper.ApiResponse{
		Data:    dataset,
		Code:    http.StatusOK,
		Message: "Successfully found by ID",
		Status:  true,
	}
}

// Store implements domain.UserService.
func (s *service) Store(user domain.User) helper.ApiResponse {
	if err := s.userRepository.Create(user); err != nil {
		return helper.ApiResponse{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed to store",
			Status:  false,
		}
	}

	return helper.ApiResponse{
		Data:    user,
		Code:    http.StatusCreated,
		Message: "Successfully stored",
		Status:  true,
	}
}

// Update implements domain.UserService.
func (s *service) Update(user domain.User) helper.ApiResponse {
	if err := s.userRepository.Update(user); err != nil {
		return helper.ApiResponse{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed to update",
			Status:  false,
		}
	}

	return helper.ApiResponse{
		Data:    user,
		Code:    http.StatusOK,
		Message: "Successfully updated",
		Status:  true,
	}
}

// FindAllPost implements domain.UserService.
func (s *service) FindAllPost(id int64) helper.ApiResponse {
	user, err := s.userRepository.FindByID(id)

	if err != nil {
		return helper.ApiResponse{
			Data:    nil,
			Code:    http.StatusNotFound,
			Message: "Failed to find by ID",
			Status:  false,
		}
	}

	dataset, err := s.userRepository.FindAllPost(id)

	if err != nil {
		return helper.ApiResponse{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed to find all post",
			Status:  false,
		}
	}

	var PostsDatas []domain.PostDataUser

	for _, post := range dataset {
		PostsDatas = append(PostsDatas, domain.PostDataUser{
			User_ID: post.User_ID,
			Title:   post.Title,
			Content: post.Content,
			Slug:    post.Slug,
		})
	}

	result := domain.UserPost{
		Username: user.Username,
		Email:    user.Email,
		Posts:    PostsDatas,
	}

	return helper.ApiResponse{
		Data:    result,
		Code:    http.StatusOK,
		Message: "Successfully found all post",
		Status:  true,
	}

}
