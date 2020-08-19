package auth

import (
	"Rihno/entity"
	"Rihno/models"
	"Rihno/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context)  {
	var registerRequest request.Register
	if err := c.ShouldBind(&registerRequest); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	var user entity.User
	user.Name = registerRequest.Name
	user.Email = registerRequest.Email
	pass, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 14)
	user.Password = string(pass)

	err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
