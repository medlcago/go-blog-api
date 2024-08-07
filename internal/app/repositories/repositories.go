package repositories

import (
	"go-blog-api/internal/app/models"
	"go-blog-api/internal/app/repositories/store/postgres"
	"go-blog-api/internal/app/utils/pagination"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(username, password string) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	FindAllUsers(p pagination.LimitOffsetPaginator) ([]models.User, int64, error)
}

type Post interface {
	CreatePost(userId uint64, title, content string) (*models.Post, error)
	FindAllPosts(pagination.LimitOffsetPaginator) ([]models.Post, int64, error)
}

type Repository struct {
	User
	Post
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: postgres.NewUserRepository(db),
		Post: postgres.NewPostRepository(db),
	}
}
