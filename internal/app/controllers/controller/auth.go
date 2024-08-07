package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/internal/app/services"
	"go-blog-api/internal/app/types"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	var createUserRequest struct {
		Username string `json:"username" binding:"ascii,required,min=5,max=20"`
		Password string `json:"password" binding:"min=6,max=64"`
	}

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, types.AppError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	user, err := ac.authService.Register(createUserRequest.Username, createUserRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.AppError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, user)

}

func (ac *AuthController) Login(c *gin.Context) {
	var userLoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, types.AppError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	token, err := ac.authService.Login(userLoginRequest.Username, userLoginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, types.AppError{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
