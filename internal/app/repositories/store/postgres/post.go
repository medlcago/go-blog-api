package postgres

import (
	"go-blog-api/internal/app/models"
	"go-blog-api/internal/app/utils/pagination"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (postRepository *PostRepository) CreatePost(userId uint64, title, content string) (*models.Post, error) {
	post := &models.Post{
		UserID:  userId,
		Title:   title,
		Content: content,
	}
	if err := postRepository.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (postRepository *PostRepository) FindAllPosts(pg pagination.LimitOffsetPaginator) ([]models.Post, int64, error) {
	var posts []models.Post
	var count int64
	if err := postRepository.db.Model(&models.Post{}).Count(&count).Preload("User").Limit(pg.GetLimit()).Offset(pg.GetOffset()).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, count, nil
}
