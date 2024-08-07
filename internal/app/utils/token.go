package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-blog-api/internal/app/config"
	"go-blog-api/internal/app/models"
	"go-blog-api/internal/app/types"
	"strconv"
	"time"
)

var (
	InvalidTokenError = errors.New("invalid token")
	BadSigningMethod  = errors.New("invalid signing method")
)

func CreateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.Itoa(int(user.ID)),
		},
	})
	return token.SignedString([]byte(config.SecretKey))
}

func ParseToken(jwtToken string) (*types.JWTClaims, error) {

	token, err := jwt.ParseWithClaims(jwtToken, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, BadSigningMethod
		}
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return &types.JWTClaims{}, InvalidTokenError
	}

	if claims, ok := token.Claims.(*types.JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return &types.JWTClaims{}, InvalidTokenError
	}
}
