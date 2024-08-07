package controllers

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/internal/app/controllers/controller"
	"go-blog-api/internal/app/services"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type PostsController interface {
	CreatePost(c *gin.Context)
	GetPosts(c *gin.Context)
	GetPostById(c *gin.Context)
}

type UsersController interface {
	GetUsers(c *gin.Context)
	GetUserById(c *gin.Context)
}

type Controller struct {
	AuthController
	PostsController
	UsersController
}

func NewController(services *services.Service) *Controller {
	return &Controller{
		AuthController:  controller.NewAuthController(services.AuthService),
		PostsController: controller.NewPostsController(services.PostService),
		UsersController: controller.NewUsersController(services.UserService),
	}
}
