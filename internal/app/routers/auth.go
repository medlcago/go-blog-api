package routers

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/internal/app/controllers"
)

func SetupAuthRoutes(r *gin.RouterGroup, authController controllers.AuthController) {
	router := r.Group("/auth")

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
}
