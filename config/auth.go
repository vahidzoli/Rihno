package config

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	UserId uint64 `json:"user_id"`
}

