package service

import (
	"go-blog-api/internal/app/models"
	"go-blog-api/internal/app/repositories"
	"go-blog-api/internal/app/types"
	"go-blog-api/internal/app/utils/pagination"
)

type PostService struct {
	postRepository repositories.Post
}

func NewPostService(postRepository repositories.Post) *PostService {
	return &PostService{
		postRepository: postRepository,
	}
}

func (p *PostService) CreatePost(userId uint64, title, content string) (*models.Post, error) {
	post, err := p.postRepository.CreatePost(userId, title, content)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostService) FetchPosts(pg pagination.LimitOffsetPaginator) (types.PaginationResponse, error) {
	posts, count, err := p.postRepository.FindAllPosts(pg)
	if err != nil {
		return types.PaginationResponse{}, err
	}
	return types.PaginationResponse{
		Count: count,
		Items: posts,
	}, nil
}

func (p *PostService) FetchPostById(id uint64) (*models.Post, error) {
	post, err := p.postRepository.FindPostById(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}
