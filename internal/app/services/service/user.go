package service

import (
	"go-blog-api/internal/app/repositories"
	"go-blog-api/internal/app/types"
	"go-blog-api/internal/app/utils/pagination"
)

type UserService struct {
	userRepository repositories.User
}

func NewUserService(userRepository repositories.User) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) FetchUsers(pg pagination.LimitOffsetPaginator) (types.PaginationResponse, error) {
	users, count, err := u.userRepository.FindAllUsers(pg)
	if err != nil {
		return types.PaginationResponse{}, err
	}
	return types.PaginationResponse{
		Count: count,
		Items: users,
	}, nil

}
