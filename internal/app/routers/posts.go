package routers

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/internal/app/controllers"
	"go-blog-api/internal/app/middlewares"
)

func SetupPostsRoutes(r *gin.RouterGroup, postController controllers.PostsController) {
	router := r.Group("/posts")

	router.POST("/", middlewares.JWTAuthMiddleware(), postController.CreatePost)
	router.GET("/", postController.GetPosts)
	router.GET("/:id", postController.GetPostById)
}
