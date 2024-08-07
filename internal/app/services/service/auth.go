package service

import (
	"errors"
	"go-blog-api/internal/app/models"
	"go-blog-api/internal/app/repositories"
	"go-blog-api/internal/app/utils"
)

var (
	InvalidCredentials = errors.New("invalid credentials")
	UserAlreadyExists  = errors.New("user already exists")
)

type AuthService struct {
	userRepository repositories.User
}

func NewAuthService(userRepository repositories.User) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (a *AuthService) Register(username string, password string) (*models.User, error) {
	_, err := a.userRepository.FindUserByUsername(username)
	if err == nil {
		return nil, UserAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := a.userRepository.CreateUser(username, hashedPassword)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (a *AuthService) Login(username string, password string) (string, error) {
	user, err := a.userRepository.FindUserByUsername(username)
	if err != nil {
		return "", InvalidCredentials
	}
	ok := utils.CheckPasswordHash(password, user.Password)
	if !ok {
		return "", InvalidCredentials
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil

}
