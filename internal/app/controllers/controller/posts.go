package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/internal/app/services"
	"go-blog-api/internal/app/types"
	"go-blog-api/internal/app/utils/pagination"
	"net/http"
	"strconv"
)

type PostsController struct {
	postService services.PostService
}

func NewPostsController(postService services.PostService) *PostsController {
	return &PostsController{
		postService: postService,
	}
}

func (pc *PostsController) CreatePost(c *gin.Context) {
	claims := c.MustGet("claims").(*types.JWTClaims)
	userId, _ := strconv.ParseUint(claims.Subject, 10, 64)

	var createPostRequest struct {
		Title   string `json:"title" binding:"min=1,max=128"`
		Content string `json:"content" binding:"min=1,max=1024"`
	}

	if err := c.ShouldBindJSON(&createPostRequest); err != nil {
		c.JSON(http.StatusBadRequest, types.AppError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	post, err := pc.postService.CreatePost(userId, createPostRequest.Title, createPostRequest.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.AppError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (pc *PostsController) GetPosts(c *gin.Context) {
	var query pagination.LimitOffset

	if err := c.ShouldBindQuery(&query); err != nil {
		query.SetDefault()
	}

	posts, err := pc.postService.FetchPosts(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.AppError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (pc *PostsController) GetPostById(c *gin.Context) {
	var postRequest struct {
		Id uint64 `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, types.AppError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	post, err := pc.postService.FetchPostById(postRequest.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, types.AppError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}
