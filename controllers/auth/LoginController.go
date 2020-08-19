package auth

import (
	"Rihno/config"
	"Rihno/entity"
	"Rihno/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var loginDto request.Login
	if err := c.ShouldBind(&loginDto); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	email := loginDto.Email
	password := loginDto.Password

	var user entity.User
	config.DB.Where("email = ?", email).First(&user)

	//compare the user from the request, with the one we defined:
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.Id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"token":token})
}

func CreateToken(userId uint64) (string, error) {
	var err error
	// Creating Access Token
	claims := &config.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(time.Hour * 15).Unix(),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(config.App.Key))
	if err != nil {
		return "", err
	}
	return token, nil
}
