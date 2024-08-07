package services

import (
	"go-blog-api/internal/app/models"
	"go-blog-api/internal/app/repositories"
	"go-blog-api/internal/app/services/service"
	"go-blog-api/internal/app/types"
	"go-blog-api/internal/app/utils/pagination"
)

type AuthService interface {
	Register(username string, password string) (*models.User, error)
	Login(username string, password string) (string, error)
}

type PostService interface {
	CreatePost(userId uint64, title, content string) (*models.Post, error)
	FetchPosts(pagination.LimitOffsetPaginator) (types.PaginationResponse, error)
	FetchPostById(uint64) (*models.Post, error)
}

type UserService interface {
	FetchUsers(p pagination.LimitOffsetPaginator) (types.PaginationResponse, error)
	FetchUserById(id uint64) (*models.User, error)
}

type Service struct {
	AuthService
	PostService
	UserService
}

func NewServices(repository *repositories.Repository) *Service {
	return &Service{
		AuthService: service.NewAuthService(repository.User),
		PostService: service.NewPostService(repository.Post),
		UserService: service.NewUserService(repository.User),
	}
}
