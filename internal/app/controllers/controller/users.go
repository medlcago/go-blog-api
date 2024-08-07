package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/internal/app/services"
	"go-blog-api/internal/app/types"
	"go-blog-api/internal/app/utils/pagination"
	"net/http"
)

type UsersController struct {
	usersService services.UserService
}

func NewUsersController(userService services.UserService) *UsersController {
	return &UsersController{
		usersService: userService,
	}
}

func (uc *UsersController) GetUsers(c *gin.Context) {
	var query pagination.LimitOffset

	if err := c.ShouldBindQuery(&query); err != nil {
		query.SetDefault()
	}

	users, err := uc.usersService.FetchUsers(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Oops, something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UsersController) GetUserById(c *gin.Context) {
	var userRequest struct {
		Id uint64 `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, types.AppError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	user, err := uc.usersService.FetchUserById(userRequest.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, types.AppError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
