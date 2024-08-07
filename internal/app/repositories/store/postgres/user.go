package postgres

import (
	"errors"
	"go-blog-api/internal/app/models"
	"go-blog-api/internal/app/utils/pagination"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (userRepository *UserRepository) CreateUser(username, password string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Password: password,
	}

	if err := userRepository.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := userRepository.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (userRepository *UserRepository) FindAllUsers(pg pagination.LimitOffsetPaginator) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	if err := userRepository.db.Model(&models.User{}).Count(&count).Limit(pg.GetLimit()).Offset(pg.GetOffset()).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (userRepository *UserRepository) FindUserById(id uint64) (*models.User, error) {
	var user models.User
	if err := userRepository.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
