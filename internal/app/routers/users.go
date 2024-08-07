package routers

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/internal/app/controllers"
)

func SetupUserRoutes(r *gin.RouterGroup, userController controllers.UsersController) {
	router := r.Group("/users")

	router.GET("/", userController.GetUsers)
	router.GET("/:id", userController.GetUserById)
}
