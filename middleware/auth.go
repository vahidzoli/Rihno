package middleware

import (
	"Rihno/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Malformed Token"))
			c.Abort()
		} else {
			jwtToken := authHeader[1]
			claims := &config.Claims{}
			token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.App.Key), nil
			})

			if token != nil && token.Valid {
				c.Set("authData", claims)
				c.Next()
			} else {
				fmt.Println(err)
				c.Writer.WriteHeader(http.StatusUnauthorized)
				c.Writer.Write([]byte("Unauthorized"))
				c.Abort()
			}
		}
	}
}

