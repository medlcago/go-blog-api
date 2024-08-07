package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5/request"
	"go-blog-api/internal/app/types"
	"go-blog-api/internal/app/utils"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := request.BearerExtractor{}.ExtractToken(c.Request)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.AppError{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		claims, err := utils.ParseToken(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.AppError{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		c.Set("claims", claims)

		c.Next()

	}
}
