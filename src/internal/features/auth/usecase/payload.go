package auth_usecase

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	jwt.RegisteredClaims
	Version int `json:"version"`
	UserId  int `json:"user_id"`
}
