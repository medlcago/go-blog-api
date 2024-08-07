package types

import "github.com/golang-jwt/jwt/v5"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PaginationResponse struct {
	Count int64       `json:"count"`
	Items interface{} `json:"items"`
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
